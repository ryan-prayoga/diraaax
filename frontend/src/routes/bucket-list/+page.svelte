<script lang="ts">
	import { onMount } from 'svelte';
	import { bucketList as bucketApi } from '$lib/api';
	import type { BucketItem } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import BucketProgressCard from '$lib/components/bucket/BucketProgressCard.svelte';
	import BucketListCard from '$lib/components/bucket/BucketListCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import ErrorAlert from '$lib/components/ui/ErrorAlert.svelte';

	let items = $state<BucketItem[]>([]);
	let loading = $state(true);
	let error = $state('');
	let creating = $state(false);
	let actionError = $state('');

	let newTitle = $state('');
	let newDescription = $state('');
	let newCategory = $state('dream');
	let newTargetDate = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const list = await bucketApi.list();
			items = Array.isArray(list) ? list : [];
		} catch (err: any) {
			error = err.message || 'Gagal memuat bucket list';
			items = [];
		} finally {
			loading = false;
		}
	}

	async function handleToggle(id: number) {
		try {
			const updated = await bucketApi.toggle(id);
			items = items.map(item => item.id === id ? updated : item);
		} catch (err: any) {
			actionError = err.message || 'Gagal mengubah status bucket list';
		}
	}

	async function handleCreateItem(e: Event) {
		e.preventDefault();
		actionError = '';
		if (!newTitle.trim()) {
			actionError = 'Judul bucket list wajib diisi ✨';
			return;
		}

		creating = true;
		try {
			const created = await bucketApi.create({
				title: newTitle.trim(),
				description: newDescription.trim() || undefined,
				category: newCategory,
				target_date: newTargetDate || undefined
			});
			items = [created, ...items];
			newTitle = '';
			newDescription = '';
			newCategory = 'dream';
			newTargetDate = '';
		} catch (err: any) {
			actionError = err.message || 'Gagal menambah bucket list';
		} finally {
			creating = false;
		}
	}

	async function handleDeleteItem(id: number) {
		actionError = '';
		try {
			await bucketApi.remove(id);
			items = items.filter(item => item.id !== id);
		} catch (err: any) {
			actionError = err.message || 'Gagal menghapus bucket list';
		}
	}

	onMount(fetchData);

	const doneItems = $derived((items ?? []).filter(i => i.is_done));
	const todoItems = $derived((items ?? []).filter(i => !i.is_done));
</script>

<svelte:head>
	<title>Bucket List — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<SectionHeader
		title="Bucket List"
		subtitle="Impian dan rencana yang ingin kita wujudkan bersama"
		emoji="✨"
	/>

	{#if loading}
		<LoadingState text="Memuat bucket list..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<div class="love-card-static p-4 md:p-5">
			<form class="space-y-3" onsubmit={handleCreateItem}>
				<p class="text-sm font-semibold text-rose-deep">Tambah impian baru 🌟</p>
				<div class="grid grid-cols-1 md:grid-cols-3 gap-3">
					<input
						type="text"
						bind:value={newTitle}
						placeholder="Judul impian"
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
					<select
						bind:value={newCategory}
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					>
						<option value="dream">Dream</option>
						<option value="travel">Travel</option>
						<option value="date">Date</option>
						<option value="home">Home</option>
					</select>
					<input
						type="date"
						bind:value={newTargetDate}
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
				</div>
				<textarea
					bind:value={newDescription}
					rows="2"
					placeholder="Detail impian (opsional)"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none resize-none"
				></textarea>
				<button type="submit" class="btn-primary text-sm px-5 py-2.5 rounded-2xl" disabled={creating}>
					{creating ? 'Menyimpan...' : 'Tambah Bucket'}
				</button>
			</form>
		</div>

		{#if actionError}
			<ErrorAlert message={actionError} />
		{/if}

		{#if items.length === 0}
			<EmptyState
				emoji="✨"
				title="Belum ada bucket list"
				subtitle="Tambahkan impian dan rencana kalian!"
			/>
		{:else}
		<!-- Progress Card -->
		<div class="animate-fade-in-up">
			<BucketProgressCard {items} />
		</div>

		<!-- Todo Items -->
		{#if todoItems.length > 0}
			<div class="space-y-3">
				<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold">Belum Tercapai 🎯</p>
				{#each todoItems as item, i}
					<div class="animate-fade-in-up" style="animation-delay: {i * 80}ms">
						<BucketListCard {item} ontoggle={handleToggle} ondelete={handleDeleteItem} />
					</div>
				{/each}
			</div>
		{/if}

		<!-- Done Items -->
		{#if doneItems.length > 0}
			<div class="space-y-3">
				<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold">Sudah Tercapai 🎉</p>
				{#each doneItems as item, i}
					<div class="animate-fade-in-up" style="animation-delay: {i * 80}ms">
						<BucketListCard {item} ontoggle={handleToggle} ondelete={handleDeleteItem} />
					</div>
				{/each}
			</div>
		{/if}
		{/if}
	{/if}
</div>

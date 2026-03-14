<script lang="ts">
	import { onMount } from 'svelte';
	import { capsules as capsulesApi } from '$lib/api';
	import type { LoveCapsule } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import CapsuleCard from '$lib/components/capsules/CapsuleCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import ErrorAlert from '$lib/components/ui/ErrorAlert.svelte';

	let allCapsules = $state<LoveCapsule[]>([]);
	let loading = $state(true);
	let error = $state('');
	let creating = $state(false);
	let actionError = $state('');

	let newTitle = $state('');
	let newMessage = $state('');
	let newOpenDate = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const list = await capsulesApi.list();
			allCapsules = Array.isArray(list) ? list : [];
		} catch (err: any) {
			error = err.message || 'Gagal memuat capsule';
			allCapsules = [];
		} finally {
			loading = false;
		}
	}

	async function handleCreateCapsule(e: Event) {
		e.preventDefault();
		actionError = '';
		if (!newTitle.trim() || !newMessage.trim() || !newOpenDate) {
			actionError = 'Judul, pesan, dan tanggal buka wajib diisi 💌';
			return;
		}

		creating = true;
		try {
			const created = await capsulesApi.create({
				title: newTitle.trim(),
				message: newMessage.trim(),
				open_date: new Date(newOpenDate).toISOString(),
				visible_to: 'both',
				theme_variant: 'romantic-pink'
			});
			allCapsules = [created, ...allCapsules];
			newTitle = '';
			newMessage = '';
			newOpenDate = '';
		} catch (err: any) {
			actionError = err.message || 'Gagal membuat capsule';
		} finally {
			creating = false;
		}
	}

	async function handleDeleteCapsule(id: number) {
		actionError = '';
		try {
			await capsulesApi.remove(id);
			allCapsules = allCapsules.filter(capsule => capsule.id !== id);
		} catch (err: any) {
			actionError = err.message || 'Gagal menghapus capsule';
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Love Capsules — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<SectionHeader
		title="Love Capsules"
		subtitle="Pesan-pesan cinta yang menunggu untuk dibuka"
		emoji="💌"
	/>

	{#if loading}
		<LoadingState text="Memuat capsule..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<div class="love-card-static p-4 md:p-5">
			<form class="space-y-3" onsubmit={handleCreateCapsule}>
				<p class="text-sm font-semibold text-rose-deep">Buat capsule baru 💌</p>
				<input
					type="text"
					bind:value={newTitle}
					placeholder="Judul capsule"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
				/>
				<textarea
					bind:value={newMessage}
					rows="3"
					placeholder="Pesan untuk masa depan"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none resize-none"
				></textarea>
				<input
					type="date"
					bind:value={newOpenDate}
					class="w-full md:w-auto px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
				/>
				<button type="submit" class="btn-primary text-sm px-5 py-2.5 rounded-2xl" disabled={creating}>
					{creating ? 'Menyimpan...' : 'Simpan Capsule'}
				</button>
			</form>
		</div>

		{#if actionError}
			<ErrorAlert message={actionError} />
		{/if}

		{#if allCapsules.length === 0}
			<EmptyState
				emoji="💌"
				title="Belum ada Love Capsule"
				subtitle="Saat kapsul pertama dibuat, surat-surat cinta kalian akan tersimpan di sini"
			/>
		{:else}
			<div class="grid grid-cols-1 lg:grid-cols-2 gap-4 lg:gap-5">
				{#each allCapsules as capsule, i}
					<div class="animate-fade-in-up" style="animation-delay: {i * 100}ms">
						<CapsuleCard {capsule} ondelete={handleDeleteCapsule} />
					</div>
				{/each}
			</div>
		{/if}
	{/if}
</div>

<script lang="ts">
	import { onMount } from 'svelte';
	import { memories as memoriesApi } from '$lib/api';
	import type { Memory } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import MemoryOfDayHero from '$lib/components/memories/MemoryOfDayHero.svelte';
	import PolaroidGrid from '$lib/components/memories/PolaroidGrid.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import ErrorAlert from '$lib/components/ui/ErrorAlert.svelte';

	let allMemories = $state<Memory[]>([]);
	let randomMemory = $state<Memory | null>(null);
	let loading = $state(true);
	let error = $state('');
	let creating = $state(false);
	let actionError = $state('');

	let newTitle = $state('');
	let newDescription = $state('');
	let newMemoryDate = $state('');
	let newImageUrl = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const [list, random] = await Promise.allSettled([
				memoriesApi.list(),
				memoriesApi.random()
			]);

			if (list.status === 'fulfilled') {
				allMemories = Array.isArray(list.value) ? list.value : [];
			}
			if (random.status === 'fulfilled') {
				randomMemory = random.value ?? null;
			}

			if (list.status === 'rejected') {
				throw new Error('Gagal memuat kenangan');
			}
		} catch (err: any) {
			error = err.message || 'Gagal memuat kenangan';
		} finally {
			loading = false;
		}
	}

	onMount(fetchData);

	async function handleCreateMemory(e: Event) {
		e.preventDefault();
		actionError = '';
		if (!newTitle.trim() && !newDescription.trim() && !newImageUrl.trim()) {
			actionError = 'Isi minimal judul, deskripsi, atau image URL ya 💗';
			return;
		}

		creating = true;
		try {
			const created = await memoriesApi.create({
				title: newTitle.trim() || undefined,
				description: newDescription.trim() || undefined,
				memory_date: newMemoryDate || undefined,
				image_url: newImageUrl.trim() || undefined
			});

			allMemories = [created, ...allMemories];
			if (!randomMemory) {
				randomMemory = created;
			}

			newTitle = '';
			newDescription = '';
			newMemoryDate = '';
			newImageUrl = '';
		} catch (err: any) {
			actionError = err.message || 'Gagal menambah kenangan';
		} finally {
			creating = false;
		}
	}

	async function handleDeleteMemory(id: number) {
		actionError = '';
		try {
			await memoriesApi.remove(id);
			allMemories = allMemories.filter(memory => memory.id !== id);
			if (randomMemory?.id === id) {
				randomMemory = allMemories.find(memory => memory.id !== id) ?? null;
			}
		} catch (err: any) {
			actionError = err.message || 'Gagal menghapus kenangan';
		}
	}
</script>

<svelte:head>
	<title>Memories — diraaax 💕</title>
</svelte:head>

<div class="space-y-6 md:space-y-8">
	<SectionHeader
		title="Our Memories"
		subtitle="Setiap momen bersama adalah harta 💎"
		emoji="📸"
	/>

	{#if loading}
		<LoadingState text="Memuat kenangan..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<div class="love-card-static p-4 md:p-5">
			<form class="space-y-3" onsubmit={handleCreateMemory}>
				<p class="text-sm font-semibold text-rose-deep">Tambah kenangan baru 📸</p>
				<div class="grid grid-cols-1 md:grid-cols-2 gap-3">
					<input
						type="text"
						bind:value={newTitle}
						placeholder="Judul kenangan (opsional)"
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
					<input
						type="date"
						bind:value={newMemoryDate}
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
				</div>
				<input
					type="url"
					bind:value={newImageUrl}
					placeholder="Image URL (opsional)"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
				/>
				<textarea
					bind:value={newDescription}
					rows="2"
					placeholder="Deskripsi kenangan (opsional)"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none resize-none"
				></textarea>
				<button type="submit" class="btn-primary text-sm px-5 py-2.5 rounded-2xl" disabled={creating}>
					{creating ? 'Menyimpan...' : 'Simpan Kenangan'}
				</button>
			</form>
		</div>

		{#if actionError}
			<ErrorAlert message={actionError} />
		{/if}

		{#if allMemories.length === 0 && !randomMemory}
			<EmptyState
				emoji="📸"
				title="Belum ada kenangan"
				subtitle="Saat foto pertama diunggah, semuanya akan muncul manis di sini 💕"
			/>
		{:else}
			<!-- Memory of the Day -->
			{#if randomMemory}
				<div class="animate-fade-in-up">
					<MemoryOfDayHero memory={randomMemory} />
				</div>
			{/if}

			<!-- All Memories Grid -->
			<div class="animate-fade-in-up">
				<PolaroidGrid memories={allMemories} ondelete={handleDeleteMemory} />
			</div>
		{/if}
	{/if}
</div>

<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { capsules as capsulesApi } from '$lib/api';
	import type { LoveCapsule, CapsuleScene } from '$lib/types';
	import { canOpenCapsule, formatDate } from '$lib/utils';
	import CapsuleSceneRenderer from '$lib/components/capsules/CapsuleSceneRenderer.svelte';
	import FloatingHearts from '$lib/components/ui/FloatingHearts.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import RomanticButton from '$lib/components/ui/RomanticButton.svelte';

	let capsule = $state<LoveCapsule | null>(null);
	let scenes = $state<CapsuleScene[]>([]);
	let loading = $state(true);
	let error = $state('');
	let opening = $state(false);
	let creatingScene = $state(false);

	let newSceneType = $state<'intro' | 'photo' | 'message' | 'quote' | 'ending'>('message');
	let newSceneTitle = $state('');
	let newSceneContent = $state('');
	let newSceneImage = $state('');

	const capsuleId = Number.parseInt($page.params.id ?? '', 10);

	async function fetchData() {
		loading = true;
		error = '';
		try {
			if (!Number.isFinite(capsuleId) || capsuleId <= 0) {
				throw new Error('ID capsule tidak valid');
			}

			const [capsuleData, sceneData] = await Promise.all([
				capsulesApi.get(capsuleId),
				capsulesApi.scenes(capsuleId)
			]);

			capsule = capsuleData;
			scenes = [...sceneData].sort((a, b) => (a.order_index ?? 0) - (b.order_index ?? 0));
		} catch (err: any) {
			error = err.message || 'Gagal memuat capsule';
		} finally {
			loading = false;
		}
	}

	async function handleOpen() {
		if (!capsule) return;
		opening = true;
		try {
			capsule = await capsulesApi.open(capsuleId);
			const sceneData = await capsulesApi.scenes(capsuleId);
			scenes = [...sceneData].sort((a, b) => (a.order_index ?? 0) - (b.order_index ?? 0));
		} catch (err: any) {
			error = err.message || 'Gagal membuka capsule';
		} finally {
			opening = false;
		}
	}

	async function handleAddScene(e: Event) {
		e.preventDefault();
		if (!capsule) return;

		creatingScene = true;
		try {
			const nextOrder = Math.max(0, ...scenes.map(scene => scene.order_index ?? scene.scene_order ?? 0)) + 1;
			const created = await capsulesApi.addScene(capsuleId, {
				scene_type: newSceneType,
				scene_order: nextOrder,
				title: newSceneTitle.trim() || undefined,
				content: newSceneContent.trim() || undefined,
				image_url: newSceneImage.trim() || undefined
			});

			scenes = [...scenes, created].sort((a, b) => (a.order_index ?? 0) - (b.order_index ?? 0));
			newSceneTitle = '';
			newSceneContent = '';
			newSceneImage = '';
		} catch (err: any) {
			error = err.message || 'Gagal menambah scene';
		} finally {
			creatingScene = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>{capsule?.title || 'Love Capsule'} — diraaax 💕</title>
</svelte:head>

{#if loading}
	<LoadingState text="Membuka kapsul cinta..." />
{:else if error}
	<ErrorState message={error} onretry={fetchData} />
{:else if capsule}
	{#if capsule.is_opened}
		<!-- Opened Capsule: Show scenes -->
		<FloatingHearts count={8} />

		<div class="relative z-10">
			<!-- Back link -->
			<div class="mb-6">
				<a href="/capsules" class="text-sm text-rose-muted hover:text-pink-500 transition-colors">
					← Kembali ke Capsules
				</a>
			</div>

			<!-- Capsule Title -->
			<div class="text-center mb-10 animate-fade-in-up">
				<div class="text-4xl mb-3">💌</div>
				<h1 class="text-3xl md:text-4xl font-extrabold bg-linear-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent">
					{capsule.title}
				</h1>
				<p class="text-rose-muted text-sm mt-2">Dibuka dengan cinta 💖</p>
			</div>

			<!-- Scenes -->
			<div class="love-card-static p-4 md:p-5 mb-6">
				<form class="space-y-3" onsubmit={handleAddScene}>
					<p class="text-sm font-semibold text-rose-deep">Tambah scene baru 🎬</p>
					<div class="grid grid-cols-1 md:grid-cols-2 gap-3">
						<select
							bind:value={newSceneType}
							class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
						>
							<option value="intro">Intro</option>
							<option value="photo">Photo</option>
							<option value="message">Message</option>
							<option value="quote">Quote</option>
							<option value="ending">Ending</option>
						</select>
						<input
							type="text"
							bind:value={newSceneTitle}
							placeholder="Judul scene (opsional)"
							class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
						/>
					</div>
					<input
						type="url"
						bind:value={newSceneImage}
						placeholder="Image URL (opsional)"
						class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
					<textarea
						bind:value={newSceneContent}
						rows="2"
						placeholder="Isi scene (opsional)"
						class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none resize-none"
					></textarea>
					<button type="submit" class="btn-primary text-sm px-5 py-2.5 rounded-2xl" disabled={creatingScene}>
						{creatingScene ? 'Menyimpan...' : 'Tambah Scene'}
					</button>
				</form>
			</div>

			{#if scenes.length > 0}
				<CapsuleSceneRenderer {scenes} />
			{:else}
				<div class="text-center py-12 animate-fade-in-up">
					{#if capsule.message}
						<div class="max-w-md mx-auto love-card-static p-6 md:p-8">
							<p class="text-rose-deep text-lg leading-relaxed whitespace-pre-line" style="font-family: var(--font-heading);">
								{capsule.message}
							</p>
							<div class="mt-4 text-right">
								<span class="text-pink-400 text-sm">with love 💕</span>
							</div>
						</div>
					{/if}
				</div>
			{/if}
		</div>
	{:else}
		<!-- Locked/Unopened Capsule -->
		<div class="min-h-[60vh] flex items-center justify-center">
			<div class="text-center max-w-sm mx-auto animate-fade-in-up">
				<div class="mb-6">
					{#if canOpenCapsule(capsule.open_date)}
						<div class="text-6xl mb-4 animate-pulse-soft">💌</div>
						<h1 class="text-2xl font-extrabold text-rose-deep mb-2">{capsule.title}</h1>
						<p class="text-rose-muted text-sm mb-6">Kapsul ini sudah siap dibuka! 🎉</p>

						<RomanticButton
							onclick={handleOpen}
							disabled={opening}
							size="lg"
						>
							{opening ? '💭 Membuka...' : '💝 Buka Kapsul Ini'}
						</RomanticButton>
					{:else}
						<div class="text-6xl mb-4">🔐</div>
						<h1 class="text-2xl font-extrabold text-rose-deep mb-2">{capsule.title}</h1>
						<p class="text-rose-muted text-sm mb-2">Kapsul ini masih terkunci</p>
						<p class="text-pink-400 text-sm font-medium">
							Dibuka pada: {formatDate(capsule.open_date, 'long')}
						</p>
						<p class="text-pink-300 text-xs mt-4">Sabar ya... akan jadi kejutan indah 💕</p>
					{/if}
				</div>

				<a href="/capsules" class="text-sm text-rose-muted hover:text-pink-500 transition-colors">
					← Kembali ke Capsules
				</a>
			</div>
		</div>
	{/if}
{/if}

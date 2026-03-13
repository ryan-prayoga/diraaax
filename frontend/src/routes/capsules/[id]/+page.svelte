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

	const capsuleId = parseInt($page.params.id);

	async function fetchData() {
		loading = true;
		error = '';
		try {
			capsule = await capsulesApi.get(capsuleId);

			if (capsule.is_opened) {
				const sceneData = await capsulesApi.scenes(capsuleId);
				scenes = sceneData.sort((a, b) => a.order_index - b.order_index);
			}
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
			scenes = sceneData.sort((a, b) => a.order_index - b.order_index);
		} catch (err: any) {
			error = err.message || 'Gagal membuka capsule';
		} finally {
			opening = false;
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
				<h1 class="text-3xl md:text-4xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent">
					{capsule.title}
				</h1>
				<p class="text-rose-muted text-sm mt-2">Dibuka dengan cinta 💖</p>
			</div>

			<!-- Scenes -->
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

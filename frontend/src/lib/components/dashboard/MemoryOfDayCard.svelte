<script lang="ts">
	import { onMount } from 'svelte';
	import { memories as memoriesApi, imageUrl } from '$lib/api';
	import type { Memory } from '$lib/types';

	let memory = $state<Memory | null>(null);
	let loading = $state(true);

	onMount(async () => {
		try {
			memory = await memoriesApi.random();
		} catch {
			// no memory available
		} finally {
			loading = false;
		}
	});
</script>

{#if loading}
	<div class="love-card-static p-5 md:p-6">
		<div class="skeleton h-4 w-32 mb-3"></div>
		<div class="skeleton h-40 w-full mb-3"></div>
		<div class="skeleton h-3 w-48"></div>
	</div>
{:else if memory}
	<div class="love-card-static overflow-hidden">
		<div class="relative">
			{#if memory.image_url}
				<img
					src={imageUrl(memory.image_url)}
					alt={memory.title}
					class="w-full h-48 object-cover"
				/>
				<div class="absolute inset-0 bg-gradient-to-t from-black/30 to-transparent"></div>
			{/if}
		</div>
		<div class="p-5">
			<p class="text-rose-muted text-xs uppercase tracking-wider mb-1">✨ Memory of the Day</p>
			<h3 class="text-lg font-bold text-rose-deep mb-1">{memory.title}</h3>
			{#if memory.description}
				<p class="text-rose-muted text-sm line-clamp-2">{memory.description}</p>
			{/if}
		</div>
	</div>
{/if}

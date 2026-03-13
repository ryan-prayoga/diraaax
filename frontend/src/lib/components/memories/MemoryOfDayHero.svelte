<script lang="ts">
	import type { Memory } from '$lib/types';
	import { imageUrl } from '$lib/api';

	let {
		memory
	}: {
		memory: Memory;
	} = $props();
</script>

<div class="love-card-static overflow-hidden relative">
	{#if memory.image_url}
		<div class="relative h-56 md:h-72">
			<img
				src={imageUrl(memory.image_url)}
				alt={memory.title}
				class="w-full h-full object-cover"
			/>
			<div class="absolute inset-0 bg-gradient-to-t from-black/50 via-transparent to-transparent"></div>
			<div class="absolute bottom-0 left-0 right-0 p-5 text-white">
				<p class="text-xs uppercase tracking-wider opacity-80 mb-1">✨ Memory of the Day</p>
				<h3 class="text-xl font-bold">{memory.title}</h3>
				{#if memory.description}
					<p class="text-sm opacity-90 mt-1 line-clamp-2">{memory.description}</p>
				{/if}
			</div>
		</div>
	{:else}
		<div class="p-6 text-center bg-gradient-to-br from-pink-50 to-pink-100">
			<span class="text-4xl block mb-3">✨</span>
			<p class="text-xs uppercase tracking-wider text-rose-muted mb-1">Memory of the Day</p>
			<h3 class="text-xl font-bold text-rose-deep">{memory.title}</h3>
			{#if memory.description}
				<p class="text-sm text-rose-muted mt-2">{memory.description}</p>
			{/if}
		</div>
	{/if}
</div>

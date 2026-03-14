<script lang="ts">
	import type { Memory } from '$lib/types';
	import { formatDate } from '$lib/utils';
	import { imageUrl } from '$lib/api';

	let {
		memory,
		rotation = 0,
		ondelete
	}: {
		memory: Memory;
		rotation?: number;
		ondelete?: (id: number) => void;
	} = $props();
</script>

<div
	class="polaroid cursor-pointer group"
	style="--rotation: {rotation}deg"
>
	{#if memory.image_url}
		<div class="relative overflow-hidden rounded-sm aspect-square">
			<img
				src={imageUrl(memory.image_url)}
				alt={memory.title}
				class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500"
			/>
		</div>
	{:else}
		<div class="aspect-square bg-pink-50 rounded-sm flex items-center justify-center">
			<span class="text-4xl">📸</span>
		</div>
	{/if}

	<div class="mt-2 text-center px-1">
		<p class="text-xs font-bold text-rose-deep truncate">{memory.title || 'Untitled memory'}</p>
		<p class="text-[10px] text-rose-muted">{formatDate(memory.memory_date || memory.created_at, 'short')}</p>
		{#if ondelete}
			<button
				type="button"
				onclick={() => ondelete?.(memory.id)}
				class="mt-1 text-[10px] text-red-400 hover:text-red-500"
			>
				Hapus
			</button>
		{/if}
	</div>
</div>

<script lang="ts">
	import type { BucketItem } from '$lib/types';
	import { formatDate } from '$lib/utils';

	let {
		item,
		ontoggle
	}: {
		item: BucketItem;
		ontoggle?: (id: number) => void;
	} = $props();
</script>

<div class="love-card-static p-4 md:p-5 {item.is_done ? 'bg-green-50/50 border-green-100!' : ''}">
	<div class="flex items-start gap-3">
		<button
			onclick={() => ontoggle?.(item.id)}
			class="flex-shrink-0 w-7 h-7 rounded-full border-2 flex items-center justify-center transition-all mt-0.5
				{item.is_done
					? 'bg-green-400 border-green-400 text-white'
					: 'border-pink-300 hover:border-pink-400 hover:bg-pink-50'}"
		>
			{#if item.is_done}
				<span class="text-xs">✓</span>
			{/if}
		</button>

		<div class="flex-1 min-w-0">
			<h3 class="font-bold text-sm {item.is_done ? 'text-green-700 line-through' : 'text-rose-deep'}">
				{item.title}
			</h3>
			{#if item.description}
				<p class="text-rose-muted text-xs mt-1 line-clamp-2">{item.description}</p>
			{/if}
			<div class="flex items-center gap-2 mt-2 flex-wrap">
				{#if item.category}
					<span class="text-[10px] font-semibold bg-pink-50 text-pink-500 px-2 py-0.5 rounded-full">
						{item.category}
					</span>
				{/if}
				{#if item.target_date}
					<span class="text-[10px] text-rose-muted">
						🎯 {formatDate(item.target_date, 'short')}
					</span>
				{/if}
			</div>
		</div>

		{#if item.is_done}
			<span class="text-xl flex-shrink-0">🎉</span>
		{/if}
	</div>
</div>

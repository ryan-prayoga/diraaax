<script lang="ts">
	import type { TimelineEvent } from '$lib/types';
	import { formatDate } from '$lib/utils';
	import { imageUrl } from '$lib/api';

	let {
		event,
		ondelete
	}: {
		event: TimelineEvent;
		ondelete?: (id: number) => void;
	} = $props();

	const typeEmojis: Record<string, string> = {
		'first_meet': '🥰',
		'anniversary': '💍',
		'date': '💕',
		'travel': '✈️',
		'milestone': '⭐',
		'special': '🌟',
		'default': '💗'
	};

	const emoji = $derived(typeEmojis[event.event_type || 'default'] || typeEmojis.default);
	const eventTypeLabel = $derived(
		(event.event_type || 'special')
			.replaceAll('_', ' ')
			.replace(/\b\w/g, (char) => char.toUpperCase())
	);
</script>

<div class="relative pl-11 md:pl-13">
	<!-- Timeline dot -->
	<div class="absolute left-[0.6rem] md:left-[0.95rem] top-5 w-5 h-5 bg-pink-400 rounded-full border-[3px] border-white shadow-md flex items-center justify-center z-10">
		<span class="text-[10px] leading-none">{emoji}</span>
	</div>

	<div class="love-card-static p-4 md:p-5 lg:p-6 bg-linear-to-br from-white to-pink-50/30">
		<div class="flex items-center justify-between gap-2 mb-3 flex-wrap">
			<span class="inline-flex items-center text-xs font-semibold text-pink-500 bg-pink-50 border border-pink-100 px-2.5 py-1 rounded-full">
				{eventTypeLabel}
			</span>
			<div class="flex items-center gap-2">
				<span class="text-[11px] md:text-xs font-semibold text-rose-muted bg-white border border-pink-100 px-2.5 py-1 rounded-full">
					{formatDate(event.event_date, 'short')}
				</span>
				{#if ondelete}
					<button
						type="button"
						onclick={() => ondelete?.(event.id)}
						class="text-[11px] md:text-xs font-semibold text-red-400 hover:text-red-500 bg-white border border-red-100 px-2.5 py-1 rounded-full"
					>
						Hapus
					</button>
				{/if}
			</div>
		</div>

		<div class="grid gap-4 {event.image_url ? 'md:grid-cols-[1fr_10rem]' : 'grid-cols-1'}">
			<div class="min-w-0">
				<h3 class="font-bold text-rose-deep text-base md:text-lg mb-2 leading-snug">{event.title}</h3>
				{#if event.description}
					<p class="text-rose-deep/75 text-sm md:text-[15px] leading-relaxed whitespace-pre-line">
						{event.description}
					</p>
				{:else}
					<p class="text-rose-muted text-sm italic">Momen manis ini belum punya deskripsi.</p>
				{/if}
			</div>

			{#if event.image_url}
				<div class="relative overflow-hidden rounded-2xl border border-pink-100 h-40 md:h-full md:min-h-34 bg-pink-50">
					<img
						src={imageUrl(event.image_url)}
						alt={event.title}
						class="w-full h-full object-cover"
					/>
					<div class="absolute inset-x-0 bottom-0 h-10 bg-linear-to-t from-black/25 to-transparent"></div>
				</div>
			{/if}
		</div>
	</div>
</div>

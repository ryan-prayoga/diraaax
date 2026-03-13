<script lang="ts">
	import type { TimelineEvent } from '$lib/types';
	import { formatDate } from '$lib/utils';
	import { imageUrl } from '$lib/api';

	let {
		event
	}: {
		event: TimelineEvent;
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
</script>

<div class="relative pl-12">
	<!-- Timeline dot -->
	<div class="absolute left-3 top-3 w-5 h-5 bg-pink-400 rounded-full border-[3px] border-white shadow-md flex items-center justify-center z-10">
		<span class="text-[10px]">{emoji}</span>
	</div>

	<div class="love-card-static p-4 md:p-5">
		{#if event.event_type}
			<span class="inline-block text-xs font-semibold text-pink-500 bg-pink-50 px-2.5 py-1 rounded-full mb-2">
				{event.event_type.replace('_', ' ')}
			</span>
		{/if}

		<div class="flex items-start justify-between gap-3">
			<div class="flex-1 min-w-0">
				<h3 class="font-bold text-rose-deep text-base mb-1">{event.title}</h3>
				<p class="text-rose-muted text-xs mb-2">{formatDate(event.event_date, 'short')}</p>
				{#if event.description}
					<p class="text-rose-deep/70 text-sm leading-relaxed">{event.description}</p>
				{/if}
			</div>
			{#if event.image_url}
				<img
					src={imageUrl(event.image_url)}
					alt={event.title}
					class="w-20 h-20 rounded-2xl object-cover flex-shrink-0"
				/>
			{/if}
		</div>
	</div>
</div>

<script lang="ts">
	import type { CapsuleScene } from '$lib/types';

	let {
		scene,
		isQuote = false
	}: {
		scene: CapsuleScene;
		isQuote?: boolean;
	} = $props();
</script>

{#if isQuote}
	<!-- Quote style -->
	<div class="max-w-md mx-auto text-center py-6 px-4">
		<div class="text-pink-300 text-6xl leading-none mb-2">"</div>
		<p class="text-xl md:text-2xl font-bold text-rose-deep italic leading-relaxed">
			{scene.content}
		</p>
		{#if scene.title}
			<p class="text-rose-muted text-sm mt-4">— {scene.title}</p>
		{/if}
	</div>
{:else}
	<!-- Letter-like message -->
	<div class="max-w-md mx-auto">
		<div class="love-card-static p-6 md:p-8 bg-gradient-to-br from-pink-50/50 to-white">
			{#if scene.title}
				<p class="text-rose-muted text-xs uppercase tracking-wider mb-3">{scene.title}</p>
			{/if}
			<p class="text-rose-deep text-base md:text-lg leading-relaxed whitespace-pre-line" style="font-family: var(--font-heading);">
				{scene.content}
			</p>
			<div class="mt-4 text-right">
				<span class="text-pink-400 text-sm">with love 💕</span>
			</div>
		</div>
	</div>
{/if}

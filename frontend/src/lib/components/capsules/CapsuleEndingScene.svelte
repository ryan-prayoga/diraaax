<script lang="ts">
	import type { CapsuleScene } from '$lib/types';

	let {
		scene
	}: {
		scene: CapsuleScene;
	} = $props();
</script>

<div class="text-center py-12 px-4">
	<div class="max-w-sm mx-auto love-card-static p-8 bg-gradient-to-br from-pink-50 to-pink-100/50">
		<div class="text-4xl mb-4">💖</div>
		{#if scene.title}
			<h3 class="text-xl font-bold text-rose-deep mb-3">{scene.title}</h3>
		{/if}
		{#if scene.content}
			<p class="text-rose-muted text-sm leading-relaxed">{scene.content}</p>
		{/if}
		<div class="mt-6 flex justify-center gap-1">
			{#each Array(5) as _}
				<span class="text-pink-300 text-sm">💗</span>
			{/each}
		</div>
		<p class="text-pink-300 text-xs mt-4">Forever & Always ∞</p>
	</div>
</div>

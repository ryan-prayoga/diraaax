<script lang="ts">
	import type { CapsuleScene } from '$lib/types';

	let {
		scene
	}: {
		scene: CapsuleScene;
	} = $props();
</script>

<div class="text-center py-8 md:py-12 px-4">
	<div class="text-5xl mb-6 animate-pulse-soft">💌</div>
	{#if scene.title}
		<h2 class="text-3xl md:text-4xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent mb-4">
			{scene.title}
		</h2>
	{/if}
	{#if scene.content}
		<p class="text-rose-muted text-lg max-w-md mx-auto leading-relaxed">
			{scene.content}
		</p>
	{/if}
	<div class="mt-6 flex justify-center gap-2">
		<span class="w-2 h-2 rounded-full bg-pink-300"></span>
		<span class="w-2 h-2 rounded-full bg-pink-400"></span>
		<span class="w-2 h-2 rounded-full bg-pink-300"></span>
	</div>
</div>

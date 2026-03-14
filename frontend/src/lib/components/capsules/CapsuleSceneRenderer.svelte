<script lang="ts">
	import type { CapsuleScene } from '$lib/types';
	import CapsuleIntroScene from './CapsuleIntroScene.svelte';
	import CapsulePhotoScene from './CapsulePhotoScene.svelte';
	import CapsuleMessageScene from './CapsuleMessageScene.svelte';
	import CapsuleEndingScene from './CapsuleEndingScene.svelte';

	let {
		scenes = []
	}: {
		scenes: CapsuleScene[];
	} = $props();
</script>

<div class="space-y-8 md:space-y-12">
	{#each scenes as scene, i}
		<div class="animate-fade-in-up" style="animation-delay: {i * 300}ms">
			{#if scene.scene_type === 'intro'}
				<CapsuleIntroScene {scene} />
			{:else if scene.scene_type === 'photo'}
				<CapsulePhotoScene {scene} />
			{:else if scene.scene_type === 'message'}
				<CapsuleMessageScene {scene} />
			{:else if scene.scene_type === 'quote'}
				<CapsuleMessageScene {scene} isQuote={true} />
			{:else if scene.scene_type === 'ending'}
				<CapsuleEndingScene {scene} />
			{:else}
				<CapsuleMessageScene {scene} />
			{/if}
		</div>
	{/each}
</div>

<script lang="ts">
	import type { CapsuleScene } from '$lib/types';
	import { imageUrl } from '$lib/api';

	let {
		scene
	}: {
		scene: CapsuleScene;
	} = $props();
</script>

<div class="max-w-lg mx-auto">
	<div class="love-card-static overflow-hidden">
		{#if scene.image_url}
			<img
				src={imageUrl(scene.image_url)}
				alt={scene.title || 'Memory'}
				class="w-full aspect-[4/3] object-cover"
			/>
		{/if}
		{#if scene.title || scene.content}
			<div class="p-5 text-center">
				{#if scene.title}
					<h3 class="font-bold text-rose-deep text-lg mb-1">{scene.title}</h3>
				{/if}
				{#if scene.content}
					<p class="text-rose-muted text-sm">{scene.content}</p>
				{/if}
			</div>
		{/if}
	</div>
</div>

<script lang="ts">
	let {
		title,
		subtitle = '',
		emoji = '',
		class: className = ''
	}: {
		title: string;
		subtitle?: string;
		emoji?: string;
		class?: string;
	} = $props();
</script>

<div class="mb-6 {className}">
	<div class="flex items-center gap-2">
		{#if emoji}
			<span class="text-2xl">{emoji}</span>
		{/if}
		<h2 class="text-2xl md:text-3xl font-bold text-rose-deep">{title}</h2>
	</div>
	{#if subtitle}
		<p class="text-rose-muted mt-1 text-sm">{subtitle}</p>
	{/if}
</div>

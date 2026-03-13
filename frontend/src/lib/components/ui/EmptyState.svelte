<script lang="ts">
	let {
		emoji = '✨',
		title = 'Belum ada data',
		subtitle = 'Data akan muncul di sini nanti',
		class: className = ''
	}: {
		emoji?: string;
		title?: string;
		subtitle?: string;
		class?: string;
	} = $props();
</script>

<div class="love-card-static p-8 md:p-12 text-center {className}">
	<div class="text-5xl mb-4 animate-pulse-soft">{emoji}</div>
	<h3 class="text-lg font-bold text-rose-deep mb-2">{title}</h3>
	<p class="text-rose-muted text-sm">{subtitle}</p>
</div>

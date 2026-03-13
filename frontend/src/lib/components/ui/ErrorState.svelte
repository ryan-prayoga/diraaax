<script lang="ts">
	let {
		message = 'Terjadi kesalahan',
		onretry = undefined as (() => void) | undefined,
		class: className = ''
	}: {
		message?: string;
		onretry?: () => void;
		class?: string;
	} = $props();
</script>

<div class="love-card-static p-6 border-red-100! text-center {className}">
	<div class="text-3xl mb-3">😿</div>
	<p class="text-red-400 font-medium mb-1">{message}</p>
	<p class="text-rose-muted text-sm mb-4">Coba lagi nanti ya...</p>
	{#if onretry}
		<button
			onclick={onretry}
			class="btn-secondary text-sm px-4 py-2"
		>
			Coba Lagi
		</button>
	{/if}
</div>

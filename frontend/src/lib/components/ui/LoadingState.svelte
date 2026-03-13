<script lang="ts">
	let {
		text = 'Memuat...',
		class: className = ''
	}: {
		text?: string;
		class?: string;
	} = $props();
</script>

<div class="flex flex-col items-center justify-center py-16 {className}">
	<div class="relative mb-4">
		<div class="w-12 h-12 rounded-full border-[3px] border-pink-200 border-t-pink-500 animate-spin"></div>
		<span class="absolute inset-0 flex items-center justify-center text-lg">💕</span>
	</div>
	<p class="text-rose-muted text-sm animate-pulse">{text}</p>
</div>

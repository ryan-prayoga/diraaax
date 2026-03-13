<script lang="ts">
	let {
		count = 8,
		class: className = ''
	}: {
		count?: number;
		class?: string;
	} = $props();

	const hearts = ['💕', '💗', '💖', '💝', '🩷', '♥️', '💘', '💓'];

	function randomStyle(i: number) {
		const left = Math.random() * 100;
		const delay = Math.random() * 8;
		const duration = 6 + Math.random() * 6;
		const size = 0.6 + Math.random() * 1;
		return `left: ${left}%; animation-delay: ${delay}s; animation-duration: ${duration}s; font-size: ${size}rem;`;
	}
</script>

<div class="fixed inset-0 pointer-events-none overflow-hidden z-0 {className}" aria-hidden="true">
	{#each Array(count) as _, i}
		<span
			class="absolute bottom-0 animate-float-heart opacity-0"
			style={randomStyle(i)}
		>
			{hearts[i % hearts.length]}
		</span>
	{/each}
</div>

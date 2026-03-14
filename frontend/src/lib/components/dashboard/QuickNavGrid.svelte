<script lang="ts">
	const navItems = [
		{ href: '/timeline', emoji: '📅', label: 'Timeline', color: 'from-pink-100 to-rose-100' },
		{ href: '/memories', emoji: '📸', label: 'Memories', color: 'from-purple-100 to-pink-100' },
		{ href: '/capsules', emoji: '💌', label: 'Love Capsules', color: 'from-red-100 to-pink-100' },
		{ href: '/bucket-list', emoji: '✨', label: 'Bucket List', color: 'from-amber-100 to-pink-100' },
		{ href: '/moods', emoji: '🥰', label: 'Moods', color: 'from-green-100 to-pink-100' },
		{ href: '/love-reasons', emoji: '💝', label: 'Love Reasons', color: 'from-fuchsia-100 to-pink-100' },
		{ href: '/our-story', emoji: '💕', label: 'Our Story', color: 'from-rose-100 to-pink-100' }
	];
</script>

<div>
	<p class="text-rose-muted text-xs uppercase tracking-wider mb-3 font-semibold">Explore</p>
	<div class="grid grid-cols-2 gap-3">
		{#each navItems as item, i}
			<a
				href={item.href}
				class="love-card p-4 text-center bg-linear-to-br {item.color} border-pink-100 hover:border-pink-200 transition-all group"
				style="animation-delay: {i * 80}ms"
			>
				<span class="text-2xl block mb-1 group-hover:scale-110 transition-transform duration-300">{item.emoji}</span>
				<span class="text-sm font-semibold text-rose-deep">{item.label}</span>
			</a>
		{/each}
	</div>
</div>

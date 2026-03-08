<script lang="ts">
	import { user } from '$lib/auth';
</script>

<svelte:head>
	<title>Dashboard - diraaax</title>
</svelte:head>

<div class="space-y-6">
	<!-- Hero Card -->
	<div class="bg-white rounded-3xl shadow-lg shadow-pink-100 p-6 border border-pink-100 text-center">
		<div class="text-4xl mb-3">💕</div>
		<h1 class="text-2xl font-bold text-pink-500 mb-1">
			Welcome{#if $user}, {$user.nickname || $user.name}{/if}!
		</h1>
		<p class="text-pink-300 text-sm">Ryan & Dira</p>
		<p class="text-pink-200 text-xs mt-1">Maboyyy & Magirll</p>
		<div class="mt-4 bg-gradient-to-r from-pink-50 to-pink-100/50 rounded-2xl p-4">
			<p class="text-pink-400 text-sm italic">
				"Every love story is beautiful, but ours is my favorite."
			</p>
		</div>
	</div>

	<!-- Quick Links -->
	<div class="grid grid-cols-2 gap-3">
		<a href="/gallery" class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 text-center hover:shadow-lg hover:border-pink-200 transition-all group">
			<div class="text-3xl mb-2 group-hover:scale-110 transition-transform">📸</div>
			<h3 class="font-semibold text-pink-500 text-sm">Our Gallery</h3>
			<p class="text-pink-300 text-xs mt-1">Photos & Memories</p>
		</a>

		<a href="/future-plans" class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 text-center hover:shadow-lg hover:border-pink-200 transition-all group">
			<div class="text-3xl mb-2 group-hover:scale-110 transition-transform">✨</div>
			<h3 class="font-semibold text-pink-500 text-sm">Future Plans</h3>
			<p class="text-pink-300 text-xs mt-1">Dreams Together</p>
		</a>

		<a href="/secret" class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 text-center hover:shadow-lg hover:border-pink-200 transition-all group">
			<div class="text-3xl mb-2 group-hover:scale-110 transition-transform">💌</div>
			<h3 class="font-semibold text-pink-500 text-sm">Secret Notes</h3>
			<p class="text-pink-300 text-xs mt-1">Just For Us</p>
		</a>

		<a href="/our-story" class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 text-center hover:shadow-lg hover:border-pink-200 transition-all group">
			<div class="text-3xl mb-2 group-hover:scale-110 transition-transform">💕</div>
			<h3 class="font-semibold text-pink-500 text-sm">Our Story</h3>
			<p class="text-pink-300 text-xs mt-1">How It Started</p>
		</a>
	</div>

	<!-- Anniversary Card -->
	<div class="bg-gradient-to-r from-pink-100 to-pink-50 rounded-3xl p-5 border border-pink-200 text-center">
		<p class="text-pink-400 text-xs uppercase tracking-wider mb-1">Anniversary</p>
		<p class="text-pink-600 font-bold text-lg">12 December 2025</p>
		<p class="text-pink-300 text-xs mt-1">Our official day 💗</p>
	</div>
</div>

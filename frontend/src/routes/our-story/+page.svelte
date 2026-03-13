<script lang="ts">
	import FloatingHearts from '$lib/components/ui/FloatingHearts.svelte';

	const storyChapters = [
		{
			emoji: '🤖',
			date: 'Oktober 2025',
			title: 'Pertemuan Pertama (di Dunia Maya)',
			content: 'Semua berawal dari sebuah bot Telegram bernama Leomatchbot. Dua orang yang tidak saling kenal, dipertemukan oleh takdir digital. Siapa sangka sebuah "match" bisa mengubah segalanya? 💫'
		},
		{
			emoji: '💌',
			date: '22 Oktober 2025',
			title: 'Ryan Kirim Love Pertama',
			content: 'Ryan adalah yang pertama mengirimkan love. Awalnya hanya sebuah keberanian kecil, tapi ternyata itu adalah langkah pertama menuju sesuatu yang sangat indah. 💕'
		},
		{
			emoji: '💬',
			date: 'Oktober - November 2025',
			title: 'Chat yang Jarang',
			content: 'Di awal, chat-chat masih jarang dan malu-malu. Pesan singkat yang kadang lama dibalas, tapi setiap notifikasi selalu bikin jantung berdebar. 🦋'
		},
		{
			emoji: '🌸',
			date: 'November 2025',
			title: 'Semakin Dekat',
			content: 'Perlahan tapi pasti, percakapan menjadi lebih dalam dan intens. Dari chat random menjadi curhat, berbagi cerita, dan saling mengenal lebih jauh. Setiap hari rasanya tidak cukup tanpa ngobrol. 💗'
		},
		{
			emoji: '🎮',
			date: 'November - Desember 2025',
			title: 'Dira Jadi Support ML Ryan',
			content: 'Dira selalu ada untuk men-support Ryan di Mobile Legends. Bukan cuma di game, tapi juga di kehidupan nyata. Ini menjadi salah satu bonding terkuat yang membuat hubungan semakin erat. 🎮💖'
		},
		{
			emoji: '🥰',
			date: '14 Desember 2025',
			title: 'Ketemu Pertama Kali!',
			content: 'Hari yang ditunggu-tunggu akhirnya datang. Pertemuan pertama secara langsung! Semua keraguan hilang, semua deg-degan terbayar. Ternyata nyata jauh lebih indah dari layar. 🌟'
		},
		{
			emoji: '💍',
			date: '12 Desember 2025',
			title: 'Tanggal Anniversary',
			content: 'Tanggal 12 Desember 2025 dipilih sebagai tanggal anniversary resmi. Hari yang menandai dimulainya babak baru — bukan lagi dua orang yang saling mengenal, tapi dua orang yang saling memilih. 💕'
		},
		{
			emoji: '💕',
			date: 'Sekarang & Selamanya',
			title: 'Maboyyy & Magirll',
			content: 'Ryan dan Dira. Maboyyy dan Magirll. Dua orang yang dipertemukan takdir, disatukan oleh cinta, dan berharap bersama selamanya. Perjalanan ini masih panjang, dan setiap harinya akan diisi dengan cinta, tawa, dan kenangan indah bersama. ∞ 💖'
		}
	];
</script>

<svelte:head>
	<title>Our Story — diraaax 💕</title>
</svelte:head>

<FloatingHearts count={6} />

<div class="relative z-10 space-y-6">
	<!-- Hero -->
	<div class="text-center py-6 md:py-10 animate-fade-in-up">
		<div class="text-5xl mb-4 animate-pulse-soft">💕</div>
		<h1 class="text-3xl md:text-4xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent mb-2">
			Our Love Story
		</h1>
		<p class="text-rose-muted text-sm">Ryan & Dira • Maboyyy & Magirll</p>
		<p class="text-pink-300 text-xs mt-1">Cerita yang dimulai dari sebuah bot Telegram 🤖💕</p>
	</div>

	<!-- Story Timeline -->
	<div class="relative">
		<!-- Vertical line -->
		<div class="absolute left-5 top-0 bottom-0 w-0.5 bg-gradient-to-b from-pink-400 via-pink-300 to-pink-100"></div>

		<div class="space-y-8">
			{#each storyChapters as chapter, i}
				<div
					class="relative pl-14 animate-fade-in-up"
					style="animation-delay: {(i + 1) * 200}ms"
				>
					<!-- Timeline dot -->
					<div class="absolute left-2.5 top-3 w-6 h-6 bg-white border-[3px] border-pink-400 rounded-full flex items-center justify-center z-10 shadow-md">
						<span class="text-xs">{chapter.emoji}</span>
					</div>

					<div class="love-card-static p-5 md:p-6">
						<p class="text-xs text-pink-400 font-semibold uppercase tracking-wider mb-1">
							{chapter.date}
						</p>
						<h3 class="text-lg font-bold text-rose-deep mb-2">
							{chapter.title}
						</h3>
						<p class="text-rose-deep/70 text-sm leading-relaxed">
							{chapter.content}
						</p>
					</div>
				</div>
			{/each}
		</div>
	</div>

	<!-- Closing -->
	<div class="text-center py-10 animate-fade-in-up" style="animation-delay: 2.2s">
		<div class="max-w-sm mx-auto love-card-static p-8 bg-gradient-to-br from-pink-50 to-pink-100/50">
			<div class="text-4xl mb-4">💖</div>
			<p class="text-rose-deep font-bold text-lg mb-2">Cerita kita belum berakhir...</p>
			<p class="text-rose-muted text-sm leading-relaxed">
				Ini baru permulaan dari seribu kenangan yang akan kita buat bersama.
				Semoga selalu bersama, selalu saling mencintai, selamanya. 💕
			</p>
			<div class="mt-4 flex justify-center gap-1">
				{#each Array(7) as _}
					<span class="text-pink-300 text-xs">💗</span>
				{/each}
			</div>
			<p class="text-pink-300 text-xs mt-3 italic">Forever & Always ∞</p>
		</div>
	</div>
</div>

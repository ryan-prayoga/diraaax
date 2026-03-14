<script lang="ts">
	let {
		nickname = 'Sayang'
	}: {
		nickname?: string;
	} = $props();

	const quotes = [
		'Setiap hari bersamamu adalah hari terbaik 💕',
		'Kamu adalah rumah yang selalu ku rindu 🏡',
		'Dalam setiap detak jantung, ada namamu 💓',
		'Bersamamu, dunia terasa lebih indah ✨',
		'Kamu alasan terindah dalam hidupku 🌸',
		'Love is you, and you are love 💖',
		'Satu-satunya tempat terbaik adalah di sisimu 💗'
	];

	const randomQuote = quotes[Math.floor(Math.random() * quotes.length)];
</script>

<div class="love-card-static p-5 md:p-6 bg-linear-to-br from-pink-50 to-pink-100/60 text-center relative overflow-hidden">
	<!-- Decorative circles -->
	<div class="absolute -top-8 -right-8 w-24 h-24 bg-pink-200/20 rounded-full"></div>
	<div class="absolute -bottom-4 -left-4 w-16 h-16 bg-pink-200/20 rounded-full"></div>

	<div class="relative z-10">
		<p class="text-pink-400 text-sm font-semibold mb-1">Welcome back,</p>
		<h1 class="text-2xl md:text-3xl font-extrabold bg-linear-to-r from-pink-600 to-pink-500 bg-clip-text text-transparent mb-2 leading-tight">
			{nickname} 💕
		</h1>
		<p class="text-pink-500 text-[11px] md:text-xs tracking-[0.16em] uppercase mb-3 font-semibold">Ryan & Dira's Universe</p>
		<p class="text-rose-text/75 text-sm italic max-w-sm mx-auto leading-relaxed">
			"{randomQuote}"
		</p>
	</div>
</div>

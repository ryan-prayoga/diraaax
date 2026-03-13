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

<div class="love-card-static p-6 md:p-8 bg-gradient-to-br from-pink-50 to-pink-100/50 text-center relative overflow-hidden">
	<!-- Decorative circles -->
	<div class="absolute -top-8 -right-8 w-24 h-24 bg-pink-200/20 rounded-full"></div>
	<div class="absolute -bottom-4 -left-4 w-16 h-16 bg-pink-200/20 rounded-full"></div>

	<div class="relative z-10">
		<p class="text-pink-300 text-sm mb-1">Welcome back,</p>
		<h1 class="text-3xl md:text-4xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent mb-2">
			{nickname} 💕
		</h1>
		<p class="text-pink-400 text-xs tracking-wider uppercase mb-4">Ryan & Dira's Universe</p>
		<p class="text-rose-muted text-sm italic max-w-xs mx-auto leading-relaxed">
			"{randomQuote}"
		</p>
	</div>
</div>

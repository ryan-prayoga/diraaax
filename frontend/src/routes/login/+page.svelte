<script lang="ts">
	import { goto } from '$app/navigation';
	import { auth } from '$lib/api';
	import { checkAuth } from '$lib/stores/auth';
	import FloatingHearts from '$lib/components/ui/FloatingHearts.svelte';

	let pin = $state('');
	let error = $state('');
	let loading = $state(false);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		const normalizedPin = pin.trim();
		if (!normalizedPin) {
			error = 'Masukkan kode akses ya 💕';
			return;
		}
		error = '';
		loading = true;
		try {
			await auth.verifyPin(normalizedPin);
			await checkAuth();
			goto('/dashboard');
		} catch (err: any) {
			error = err.code === 'invalid_pin' || err.message === 'invalid_pin'
				? 'Kode akses salah, coba lagi ya 🥺'
				: (err.message || 'Terjadi kesalahan');
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>Login — diraaax 💕</title>
</svelte:head>

<FloatingHearts count={12} />

<div class="min-h-[85vh] flex items-center justify-center px-4 relative z-10">
	<div class="w-full max-w-sm">
		<!-- Login Card -->
		<div class="love-card-static p-8 md:p-10 animate-fade-in-up bg-white/90 backdrop-blur-sm">
			<div class="text-center mb-8">
				<div class="text-4xl mb-4 animate-pulse-soft">💕</div>
				<h1 class="text-4xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent mb-2">
					diraaax
				</h1>
				<p class="text-rose-muted text-sm font-medium">Ryan & Dira's Private Space</p>
				<p class="text-pink-300 text-xs mt-1">Maboyyy & Magirll</p>
			</div>

			<form onsubmit={handleSubmit} class="space-y-5">
				<div>
					<label for="pin" class="block text-sm font-semibold text-rose-muted mb-2">
						Access Code
					</label>
					<input
						id="pin"
						type="password"
						bind:value={pin}
						placeholder="Masukkan kode rahasia..."
						class="w-full px-5 py-3.5 rounded-2xl border-2 border-pink-100 focus:border-pink-400 focus:outline-none text-center text-lg tracking-[0.3em] bg-pink-50/30 placeholder:text-pink-200 placeholder:tracking-normal placeholder:text-sm transition-all duration-300 focus:shadow-[0_0_20px_rgba(236,72,153,0.15)]"
						autocomplete="off"
					/>
				</div>

				{#if error}
					<div class="text-red-400 text-sm text-center bg-red-50 rounded-2xl p-3 animate-fade-in border border-red-100">
						{error}
					</div>
				{/if}

				<button
					type="submit"
					disabled={loading}
					class="w-full btn-primary py-4 text-base rounded-2xl"
				>
					{loading ? '💭 Checking...' : 'Enter Our World 💕'}
				</button>
			</form>
		</div>

		<p class="text-center text-pink-300 text-xs mt-8 animate-fade-in" style="animation-delay: 0.5s">
			Made with love, for us 💗
		</p>
	</div>
</div>

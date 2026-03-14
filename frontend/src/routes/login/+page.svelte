<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { auth } from '$lib/api';
	import { checkAuth } from '$lib/stores/auth';
	import FloatingHearts from '$lib/components/ui/FloatingHearts.svelte';

	const PIN_LENGTH = 6;
	const keypadNumbers = ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0'];

	let pin = $state('');
	let error = $state('');
	let loading = $state(false);

	const canSubmit = $derived(pin.length === PIN_LENGTH && !loading);

	function sanitizePin(value: string) {
		return value.replace(/\D/g, '').slice(0, PIN_LENGTH);
	}

	function handlePinInput(event: Event) {
		const target = event.currentTarget as HTMLInputElement;
		pin = sanitizePin(target.value);
		error = '';
	}

	function handleDigitTap(digit: string) {
		if (loading || pin.length >= PIN_LENGTH) return;
		pin = `${pin}${digit}`;
		error = '';
		if (pin.length === PIN_LENGTH) {
			void submitPin();
		}
	}

	function handleBackspace() {
		if (loading || pin.length === 0) return;
		pin = pin.slice(0, -1);
		error = '';
	}

	function handleClear() {
		if (loading || pin.length === 0) return;
		pin = '';
		error = '';
	}

	async function submitPin() {
		const normalizedPin = pin.trim();
		if (!normalizedPin) {
			error = 'Masukkan kode akses ya 💕';
			return false;
		}
		if (normalizedPin.length < PIN_LENGTH) {
			error = `PIN harus ${PIN_LENGTH} digit ya 💗`;
			return false;
		}
		error = '';
		loading = true;
		try {
			await auth.verifyPin(normalizedPin);
			await checkAuth();
			await goto('/dashboard');
			return true;
		} catch (err: any) {
			error = err.code === 'invalid_pin' || err.message === 'invalid_pin' || err.message === 'invalid pin'
				? 'Kode akses salah, coba lagi ya 🥺'
				: (err.message || 'Terjadi kesalahan');
			pin = '';
			return false;
		} finally {
			loading = false;
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		await submitPin();
	}

	onMount(() => {
		const onKeyDown = (event: KeyboardEvent) => {
			if (loading || event.metaKey || event.ctrlKey || event.altKey) return;

			if (/^[0-9]$/.test(event.key)) {
				event.preventDefault();
				handleDigitTap(event.key);
				return;
			}

			if (event.key === 'Backspace') {
				event.preventDefault();
				handleBackspace();
				return;
			}

			if (event.key === 'Escape') {
				event.preventDefault();
				handleClear();
				return;
			}

			if (event.key === 'Enter' && pin.length === PIN_LENGTH) {
				event.preventDefault();
				void submitPin();
			}
		};

		window.addEventListener('keydown', onKeyDown);
		return () => window.removeEventListener('keydown', onKeyDown);
	});
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
				<h1 class="text-4xl font-extrabold bg-linear-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent mb-2">
					diraaax
				</h1>
				<p class="text-rose-muted text-sm font-medium">Ryan & Dira's Private Space</p>
				<p class="text-pink-300 text-xs mt-1">Maboyyy & Magirll</p>
			</div>

			<form onsubmit={handleSubmit} class="space-y-5">
				<div>
					<label for="pin-sr-only" class="block text-sm font-semibold text-rose-muted mb-2 text-center">
						Enter 6-digit Access PIN
					</label>
					<input
						id="pin-sr-only"
						type="text"
						value={pin}
						oninput={handlePinInput}
						readonly
						tabindex="-1"
						inputmode="numeric"
						autocomplete="off"
						pattern="[0-9]*"
						maxlength={PIN_LENGTH}
						aria-label="PIN akses {PIN_LENGTH} digit"
						class="sr-only"
					/>

					<div class="w-full bg-pink-50/70 border-2 border-pink-100 rounded-3xl px-4 py-4" role="group" aria-label="PIN slots">
						<div class="flex items-center justify-center gap-2.5">
							{#each Array(PIN_LENGTH) as _, idx}
								<div
									class="h-3.5 w-3.5 rounded-full border transition-all duration-200 {idx < pin.length
										? 'bg-pink-500 border-pink-500 scale-105'
										: 'bg-white border-pink-200'}"
								></div>
							{/each}
						</div>
					</div>
					<p class="text-center text-[11px] text-rose-muted mt-2" aria-live="polite">
						{pin.length}/{PIN_LENGTH} digit • Gunakan keypad atau keyboard angka
					</p>

					<div class="grid grid-cols-3 gap-3 pt-1">
						{#each keypadNumbers as digit}
							{#if digit === '0'}
								<div></div>
							{/if}
							<button
								type="button"
								onclick={() => handleDigitTap(digit)}
								disabled={loading || pin.length >= PIN_LENGTH}
								class="h-14 md:h-15 rounded-2xl border border-pink-200 bg-pink-100/85 text-rose-deep text-xl font-bold active:scale-[0.98] hover:bg-pink-200/80 transition disabled:opacity-50 disabled:cursor-not-allowed"
								aria-label={`Digit ${digit}`}
							>
								{digit}
							</button>
							{#if digit === '0'}
								<button
									type="button"
									onclick={handleBackspace}
									disabled={loading || pin.length === 0}
									class="h-14 md:h-15 rounded-2xl border border-pink-200 bg-white text-pink-500 text-base font-semibold active:scale-[0.98] hover:bg-pink-50 transition disabled:opacity-50 disabled:cursor-not-allowed"
									aria-label="Hapus digit terakhir"
								>
									⌫
								</button>
							{/if}
						{/each}
					</div>

					<div class="flex justify-center">
						<button
							type="button"
							onclick={handleClear}
							disabled={loading || pin.length === 0}
							class="text-xs text-pink-500 hover:text-pink-600 disabled:opacity-40"
						>
							Clear PIN
						</button>
					</div>
				</div>

				{#if error}
					<div class="text-red-400 text-sm text-center bg-red-50 rounded-2xl p-3 animate-fade-in border border-red-100">
						{error}
					</div>
				{/if}

				<button
					type="submit"
					disabled={!canSubmit}
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

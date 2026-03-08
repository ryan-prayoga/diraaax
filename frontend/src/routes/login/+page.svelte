<script lang="ts">
	import { goto } from '$app/navigation';
	import { api } from '$lib/api';
	import { checkAuth } from '$lib/auth';

	let pin = $state('');
	let error = $state('');
	let loading = $state(false);

	async function handleSubmit(e: Event) {
		e.preventDefault();
		const normalizedPin = pin.trim();
		if (!normalizedPin) {
			error = 'Please enter the access code';
			return;
		}
		error = '';
		loading = true;
		try {
			await api.verifyPin(normalizedPin);
			await checkAuth();
			goto('/dashboard');
		} catch (err: any) {
			error = err.message === 'invalid_pin' ? 'Invalid access code' : (err.message || 'Something went wrong');
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>Login - diraaax</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center px-4">
	<div class="w-full max-w-sm">
		<div class="bg-white rounded-3xl shadow-lg shadow-pink-100 p-8 border border-pink-100">
			<div class="text-center mb-8">
				<h1 class="text-3xl font-bold text-pink-500 mb-2">diraaax</h1>
				<p class="text-pink-300 text-sm">Ryan & Dira's Private Space</p>
				<p class="text-pink-200 text-xs mt-1">Maboyyy & Magirll</p>
			</div>

			<form onsubmit={handleSubmit} class="space-y-4">
				<div>
					<label for="pin" class="block text-sm font-medium text-pink-400 mb-2">
						Access Code
					</label>
					<input
						id="pin"
						type="password"
						bind:value={pin}
						placeholder="Enter your secret code..."
						class="w-full px-4 py-3 rounded-2xl border-2 border-pink-100 focus:border-pink-300 focus:outline-none text-center text-lg tracking-widest bg-pink-50/50 placeholder:text-pink-200 placeholder:tracking-normal placeholder:text-sm transition-colors"
						autocomplete="off"
					/>
				</div>

				{#if error}
					<div class="text-red-400 text-sm text-center bg-red-50 rounded-xl p-2">
						{error}
					</div>
				{/if}

				<button
					type="submit"
					disabled={loading}
					class="w-full bg-pink-400 hover:bg-pink-500 disabled:bg-pink-200 text-white font-medium py-3 rounded-2xl transition-colors shadow-md shadow-pink-200"
				>
					{loading ? 'Checking...' : 'Enter Our World 💕'}
				</button>
			</form>
		</div>

		<p class="text-center text-pink-200 text-xs mt-6">
			Made with love for us
		</p>
	</div>
</div>

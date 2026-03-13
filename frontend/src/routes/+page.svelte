<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { isAuthenticated, authLoading } from '$lib/stores/auth';

	onMount(() => {
		const unsub = authLoading.subscribe(loading => {
			if (!loading) {
				if ($isAuthenticated) {
					goto('/dashboard');
				} else {
					goto('/login');
				}
				unsub();
			}
		});
	});
</script>

<div class="min-h-[80vh] flex items-center justify-center">
	<div class="text-center">
		<div class="text-4xl mb-3 animate-pulse-soft">💕</div>
		<p class="text-rose-muted text-sm animate-pulse">Loading...</p>
	</div>
</div>

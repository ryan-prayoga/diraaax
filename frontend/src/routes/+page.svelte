<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { isAuthenticated, authLoading } from '$lib/auth';

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

<div class="min-h-screen flex items-center justify-center">
	<div class="text-pink-400 text-xl animate-pulse">Loading...</div>
</div>

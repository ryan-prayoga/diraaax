<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { checkAuth, isAuthenticated, user, authLoading } from '$lib/auth';
	import { api } from '$lib/api';

	let { children } = $props();

	const publicRoutes = ['/login'];

	onMount(async () => {
		await checkAuth();
	});

	$effect(() => {
		if (!$authLoading) {
			const currentPath = $page.url.pathname;
			if (!$isAuthenticated && !publicRoutes.includes(currentPath)) {
				goto('/login');
			}
			if ($isAuthenticated && currentPath === '/login') {
				goto('/dashboard');
			}
		}
	});

	async function handleLogout() {
		try {
			await api.logout();
		} catch {}
		isAuthenticated.set(false);
		user.set(null);
		goto('/login');
	}
</script>

{#if $authLoading}
	<div class="min-h-screen flex items-center justify-center">
		<div class="text-pink-400 text-xl animate-pulse">Loading...</div>
	</div>
{:else}
	{#if $isAuthenticated && $page.url.pathname !== '/login'}
		<nav class="bg-white/80 backdrop-blur-sm border-b border-pink-100 sticky top-0 z-50">
			<div class="max-w-lg mx-auto px-4 py-3 flex items-center justify-between">
				<a href="/dashboard" class="text-xl font-bold text-pink-500 tracking-tight">diraaax</a>
				<div class="flex items-center gap-3">
					{#if $user}
						<span class="text-sm text-pink-400">{$user.nickname || $user.name}</span>
					{/if}
					<button
						onclick={handleLogout}
						class="text-xs bg-pink-100 hover:bg-pink-200 text-pink-600 px-3 py-1.5 rounded-full transition-colors"
					>
						Logout
					</button>
				</div>
			</div>
		</nav>
	{/if}

	<main class="max-w-lg mx-auto px-4 py-6">
		{@render children()}
	</main>

	{#if $isAuthenticated && $page.url.pathname !== '/login'}
		<nav class="fixed bottom-0 left-0 right-0 bg-white/90 backdrop-blur-sm border-t border-pink-100 z-50">
			<div class="max-w-lg mx-auto px-2 py-2 flex justify-around">
				<a href="/dashboard" class="flex flex-col items-center gap-0.5 text-xs px-2 py-1 rounded-lg {$page.url.pathname === '/dashboard' ? 'text-pink-600 bg-pink-50' : 'text-pink-400 hover:text-pink-500'}">
					<span class="text-lg">🏠</span>
					<span>Home</span>
				</a>
				<a href="/gallery" class="flex flex-col items-center gap-0.5 text-xs px-2 py-1 rounded-lg {$page.url.pathname === '/gallery' ? 'text-pink-600 bg-pink-50' : 'text-pink-400 hover:text-pink-500'}">
					<span class="text-lg">📸</span>
					<span>Gallery</span>
				</a>
				<a href="/future-plans" class="flex flex-col items-center gap-0.5 text-xs px-2 py-1 rounded-lg {$page.url.pathname === '/future-plans' ? 'text-pink-600 bg-pink-50' : 'text-pink-400 hover:text-pink-500'}">
					<span class="text-lg">✨</span>
					<span>Plans</span>
				</a>
				<a href="/secret" class="flex flex-col items-center gap-0.5 text-xs px-2 py-1 rounded-lg {$page.url.pathname === '/secret' ? 'text-pink-600 bg-pink-50' : 'text-pink-400 hover:text-pink-500'}">
					<span class="text-lg">💌</span>
					<span>Secret</span>
				</a>
				<a href="/our-story" class="flex flex-col items-center gap-0.5 text-xs px-2 py-1 rounded-lg {$page.url.pathname === '/our-story' ? 'text-pink-600 bg-pink-50' : 'text-pink-400 hover:text-pink-500'}">
					<span class="text-lg">💕</span>
					<span>Story</span>
				</a>
			</div>
		</nav>
		<div class="h-16"></div>
	{/if}
{/if}

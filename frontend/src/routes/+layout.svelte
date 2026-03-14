<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { checkAuth, isAuthenticated, user, authLoading, logout } from '$lib/stores/auth';

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
		await logout();
		goto('/login');
	}

	const bottomNavLinks = [
		{ href: '/dashboard', emoji: '🏠', label: 'Home' },
		{ href: '/timeline', emoji: '📅', label: 'Timeline' },
		{ href: '/memories', emoji: '📸', label: 'Memories' },
		{ href: '/capsules', emoji: '💌', label: 'Capsules' },
		{ href: '/our-story', emoji: '💕', label: 'Story' }
	];

	function isActive(currentPath: string, href: string): boolean {
		return currentPath === href || currentPath.startsWith(href + '/');
	}
</script>

{#if $authLoading}
	<div class="min-h-screen flex items-center justify-center bg-rose-bg">
		<div class="text-center">
			<div class="text-4xl mb-3 animate-pulse-soft">💕</div>
			<p class="text-rose-muted text-sm animate-pulse">Loading...</p>
		</div>
	</div>
{:else}
	{#if $isAuthenticated && !publicRoutes.includes($page.url.pathname)}
		<!-- Top Navigation Bar -->
		<header class="bg-white/85 backdrop-blur-lg border-b border-pink-100 sticky top-0 z-50">
			<div class="page-container py-4 md:py-4 min-h-12 md:min-h-14 flex items-center justify-between gap-2">
				<a href="/dashboard" class="flex items-center gap-2">
					<span class="text-2xl font-extrabold bg-linear-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent tracking-tight leading-none">
						diraaax
					</span>
					<span class="text-xs text-pink-300 hidden sm:inline">💕</span>
				</a>

				<div class="flex items-center gap-3">
					{#if $user}
						<span class="text-sm text-rose-muted font-medium hidden sm:inline">
							{$user.nickname || $user.display_name}
						</span>
					{/if}

					<button
						onclick={handleLogout}
						class="text-xs bg-pink-50 hover:bg-pink-100 text-pink-500 px-3 py-1.5 rounded-full transition-colors font-medium border border-pink-100"
					>
						Logout
					</button>
				</div>
			</div>
		</header>
		<div class="h-4 md:h-6"></div>
	{/if}

	<!-- Main Content -->
	<main
		class="page-container pb-6 md:pb-8"
		class:has-mobile-bottom-nav={$isAuthenticated && !publicRoutes.includes($page.url.pathname)}
	>
		{@render children()}
	</main>

	{#if $isAuthenticated && !publicRoutes.includes($page.url.pathname)}
		<!-- Bottom Navigation (Mobile) -->
		<nav class="mobile-bottom-nav fixed bottom-0 left-0 right-0 bg-white/90 backdrop-blur-lg border-t border-pink-100 z-50 sm:hidden">
			<div class="flex justify-around px-2 py-2">
				{#each bottomNavLinks as link}
					<a
						href={link.href}
						class="flex flex-col items-center gap-0.5 text-[10px] px-2 py-1.5 rounded-xl transition-colors min-w-12
							{isActive($page.url.pathname, link.href) ? 'text-pink-600 bg-pink-50 font-semibold' : 'text-rose-muted hover:text-pink-400'}"
					>
						<span class="text-lg">{link.emoji}</span>
						<span>{link.label}</span>
					</a>
				{/each}
			</div>
		</nav>
	{/if}
{/if}

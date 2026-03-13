<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { checkAuth, isAuthenticated, user, authLoading, logout } from '$lib/stores/auth';

	let { children } = $props();

	let menuOpen = $state(false);

	const publicRoutes = ['/login'];

	onMount(async () => {
		await checkAuth();
	});

	$effect(() => {
		if (!$authLoading) {
			const currentPath = page.url.pathname;
			if (!$isAuthenticated && !publicRoutes.includes(currentPath)) {
				goto('/login');
			}
			if ($isAuthenticated && currentPath === '/login') {
				goto('/dashboard');
			}
		}
	});

	// Close menu on route change
	$effect(() => {
		page.url.pathname;
		menuOpen = false;
	});

	async function handleLogout() {
		await logout();
		goto('/login');
	}

	const navLinks = [
		{ href: '/dashboard', emoji: '🏠', label: 'Home' },
		{ href: '/timeline', emoji: '📅', label: 'Timeline' },
		{ href: '/memories', emoji: '📸', label: 'Memories' },
		{ href: '/capsules', emoji: '💌', label: 'Capsules' },
		{ href: '/bucket-list', emoji: '✨', label: 'Bucket List' },
		{ href: '/moods', emoji: '🥰', label: 'Moods' },
		{ href: '/love-reasons', emoji: '💝', label: 'Reasons' },
		{ href: '/our-story', emoji: '💕', label: 'Story' }
	];

	const bottomNavLinks = [
		{ href: '/dashboard', emoji: '🏠', label: 'Home' },
		{ href: '/timeline', emoji: '📅', label: 'Timeline' },
		{ href: '/memories', emoji: '📸', label: 'Memories' },
		{ href: '/capsules', emoji: '💌', label: 'Capsules' },
		{ href: '/our-story', emoji: '💕', label: 'Story' }
	];

	function isActive(href: string): boolean {
		return page.url.pathname === href || page.url.pathname.startsWith(href + '/');
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
	{#if $isAuthenticated && !publicRoutes.includes(page.url.pathname)}
		<!-- Top Navigation Bar -->
		<nav class="bg-white/85 backdrop-blur-lg border-b border-pink-100 sticky top-0 z-50">
			<div class="page-container py-3 flex items-center justify-between">
				<a href="/dashboard" class="flex items-center gap-2">
					<span class="text-xl font-extrabold bg-gradient-to-r from-pink-500 to-pink-400 bg-clip-text text-transparent tracking-tight">
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

					<!-- Mobile menu toggle -->
					<button
						onclick={() => menuOpen = !menuOpen}
						class="sm:hidden text-pink-400 hover:text-pink-500 transition-colors p-1"
					>
						{#if menuOpen}
							<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
							</svg>
						{:else}
							<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
							</svg>
						{/if}
					</button>

					<button
						onclick={handleLogout}
						class="text-xs bg-pink-50 hover:bg-pink-100 text-pink-500 px-3 py-1.5 rounded-full transition-colors font-medium hidden sm:inline-block"
					>
						Logout
					</button>
				</div>
			</div>

			<!-- Mobile dropdown menu -->
			{#if menuOpen}
				<div class="sm:hidden border-t border-pink-50 bg-white/95 backdrop-blur-lg animate-fade-in">
					<div class="page-container py-3 space-y-1">
						{#each navLinks as link}
							<a
								href={link.href}
								class="flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm transition-colors
									{isActive(link.href) ? 'text-pink-600 bg-pink-50 font-semibold' : 'text-rose-muted hover:bg-pink-50/50'}"
							>
								<span>{link.emoji}</span>
								<span>{link.label}</span>
							</a>
						{/each}
						<button
							onclick={handleLogout}
							class="flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm text-red-400 hover:bg-red-50/50 w-full text-left mt-2"
						>
							<span>👋</span>
							<span>Logout</span>
						</button>
					</div>
				</div>
			{/if}
		</nav>
	{/if}

	<!-- Main Content -->
	<main class="page-container py-6 md:py-8 {$isAuthenticated && !publicRoutes.includes(page.url.pathname) ? 'pb-24 sm:pb-8' : ''}">
		{@render children()}
	</main>

	{#if $isAuthenticated && !publicRoutes.includes(page.url.pathname)}
		<!-- Bottom Navigation (Mobile) -->
		<nav class="fixed bottom-0 left-0 right-0 bg-white/90 backdrop-blur-lg border-t border-pink-100 z-50 sm:hidden">
			<div class="flex justify-around px-2 py-2">
				{#each bottomNavLinks as link}
					<a
						href={link.href}
						class="flex flex-col items-center gap-0.5 text-[10px] px-2 py-1.5 rounded-xl transition-colors min-w-[3rem]
							{isActive(link.href) ? 'text-pink-600 bg-pink-50 font-semibold' : 'text-rose-muted hover:text-pink-400'}"
					>
						<span class="text-lg">{link.emoji}</span>
						<span>{link.label}</span>
					</a>
				{/each}
			</div>
		</nav>
	{/if}
{/if}

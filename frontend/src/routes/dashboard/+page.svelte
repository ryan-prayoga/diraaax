<script lang="ts">
	import { user } from '$lib/stores/auth';
	import WelcomeHero from '$lib/components/dashboard/WelcomeHero.svelte';
	import AnniversaryCounter from '$lib/components/dashboard/AnniversaryCounter.svelte';
	import MemoryOfDayCard from '$lib/components/dashboard/MemoryOfDayCard.svelte';
	import QuickNavGrid from '$lib/components/dashboard/QuickNavGrid.svelte';
</script>

<svelte:head>
	<title>Dashboard — diraaax 💕</title>
</svelte:head>

<div class="space-y-5 md:space-y-6">
	<!-- Welcome Hero -->
	<div class="animate-fade-in-up">
		<WelcomeHero nickname={$user?.nickname || $user?.display_name || 'Sayang'} />
	</div>

	<!-- Two column on desktop -->
	<div class="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-5">
		<!-- Anniversary Counter -->
		<div class="animate-fade-in-up" style="animation-delay: 150ms">
			<AnniversaryCounter />
		</div>

		<!-- Memory of the Day -->
		<div class="animate-fade-in-up" style="animation-delay: 250ms">
			<MemoryOfDayCard />
		</div>
	</div>

	<!-- Quick Navigation Grid -->
	<div class="animate-fade-in-up" style="animation-delay: 350ms">
		<QuickNavGrid />
	</div>
</div>

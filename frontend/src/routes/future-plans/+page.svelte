<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';

	let plans: any[] = $state([]);
	let loading = $state(true);
	let showForm = $state(false);
	let title = $state('');
	let description = $state('');
	let category = $state('date');
	let submitting = $state(false);

	onMount(async () => {
		await loadPlans();
	});

	async function loadPlans() {
		loading = true;
		try {
			plans = await api.getPlans();
		} catch {
			plans = [];
		} finally {
			loading = false;
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!title.trim()) return;
		submitting = true;
		try {
			await api.createPlan({
				title: title.trim(),
				description: description.trim() || undefined,
				category: category || undefined
			});
			title = '';
			description = '';
			category = 'date';
			showForm = false;
			await loadPlans();
		} catch {
		} finally {
			submitting = false;
		}
	}

	async function togglePlan(id: number) {
		try {
			await api.togglePlan(id);
			await loadPlans();
		} catch {}
	}

	async function deletePlan(id: number) {
		if (!confirm('Delete this plan?')) return;
		try {
			await api.deletePlan(id);
			await loadPlans();
		} catch {}
	}

	const categories: Record<string, string> = {
		date: '🌹 Date',
		travel: '✈️ Travel',
		milestone: '🎯 Milestone',
		food: '🍕 Food',
		activity: '🎮 Activity',
		other: '💫 Other'
	};
</script>

<svelte:head>
	<title>Future Plans - diraaax</title>
</svelte:head>

<div class="space-y-4">
	<div class="flex items-center justify-between">
		<h1 class="text-xl font-bold text-pink-500">✨ Our Future Plans</h1>
		<button
			onclick={() => showForm = !showForm}
			class="bg-pink-400 hover:bg-pink-500 text-white text-sm px-4 py-2 rounded-full transition-colors shadow-md"
		>
			{showForm ? 'Cancel' : '+ Add'}
		</button>
	</div>

	{#if showForm}
		<form onsubmit={handleSubmit} class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 space-y-3">
			<input
				type="text"
				bind:value={title}
				placeholder="What's the plan?"
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
			/>
			<textarea
				bind:value={description}
				placeholder="Tell me more... (optional)"
				rows="2"
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30 resize-none"
			></textarea>
			<select
				bind:value={category}
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
			>
				{#each Object.entries(categories) as [value, label]}
					<option {value}>{label}</option>
				{/each}
			</select>
			<button
				type="submit"
				disabled={submitting || !title.trim()}
				class="w-full bg-pink-400 hover:bg-pink-500 disabled:bg-pink-200 text-white py-2.5 rounded-xl transition-colors text-sm font-medium"
			>
				{submitting ? 'Adding...' : 'Add Plan 💕'}
			</button>
		</form>
	{/if}

	{#if loading}
		<div class="text-center text-pink-300 py-8 animate-pulse">Loading plans...</div>
	{:else if plans.length === 0}
		<div class="bg-white rounded-2xl shadow-md shadow-pink-50 p-8 border border-pink-100 text-center">
			<div class="text-4xl mb-3">🌟</div>
			<p class="text-pink-400">No plans yet!</p>
			<p class="text-pink-300 text-sm mt-1">Start dreaming together</p>
		</div>
	{:else}
		<div class="space-y-3">
			{#each plans as plan}
				<div class="bg-white rounded-2xl shadow-md shadow-pink-50 p-4 border border-pink-100 {plan.status === 'done' ? 'opacity-70' : ''}">
					<div class="flex items-start gap-3">
						<button
							onclick={() => togglePlan(plan.id)}
							class="mt-0.5 w-6 h-6 rounded-full border-2 flex items-center justify-center flex-shrink-0 transition-colors {plan.status === 'done' ? 'bg-pink-400 border-pink-400 text-white' : 'border-pink-200 hover:border-pink-400'}"
						>
							{#if plan.status === 'done'}
								<span class="text-xs">✓</span>
							{/if}
						</button>
						<div class="flex-1 min-w-0">
							<h3 class="font-medium text-pink-600 text-sm {plan.status === 'done' ? 'line-through' : ''}">{plan.title}</h3>
							{#if plan.description}
								<p class="text-pink-300 text-xs mt-1">{plan.description}</p>
							{/if}
							<div class="flex items-center gap-2 mt-2">
								{#if plan.category}
									<span class="text-xs bg-pink-50 text-pink-400 px-2 py-0.5 rounded-full">{categories[plan.category] || plan.category}</span>
								{/if}
								{#if plan.status === 'done' && plan.completed_at}
									<span class="text-xs text-pink-300">Completed {new Date(plan.completed_at).toLocaleDateString()}</span>
								{/if}
							</div>
						</div>
						<button
							onclick={() => deletePlan(plan.id)}
							class="text-pink-200 hover:text-red-400 transition-colors text-sm flex-shrink-0"
						>
							✕
						</button>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>

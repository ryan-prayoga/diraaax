<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';

	let notes: any[] = $state([]);
	let loading = $state(true);
	let showForm = $state(false);
	let title = $state('');
	let content = $state('');
	let noteType = $state('letter');
	let visibleTo = $state('both');
	let submitting = $state(false);

	onMount(async () => {
		await loadNotes();
	});

	async function loadNotes() {
		loading = true;
		try {
			notes = await api.getNotes();
		} catch {
			notes = [];
		} finally {
			loading = false;
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!title.trim() || !content.trim()) return;
		submitting = true;
		try {
			await api.createNote({
				title: title.trim(),
				content: content.trim(),
				note_type: noteType || undefined,
				visible_to: visibleTo || undefined
			});
			title = '';
			content = '';
			noteType = 'letter';
			visibleTo = 'both';
			showForm = false;
			await loadNotes();
		} catch {
		} finally {
			submitting = false;
		}
	}

	const typeLabels: Record<string, string> = {
		letter: '💌 Love Letter',
		note: '📝 Note',
		wish: '🌟 Wish',
		memory: '💭 Memory',
		promise: '🤞 Promise'
	};
</script>

<svelte:head>
	<title>Secret Notes - diraaax</title>
</svelte:head>

<div class="space-y-4">
	<div class="flex items-center justify-between">
		<h1 class="text-xl font-bold text-pink-500">💌 Secret Notes</h1>
		<button
			onclick={() => showForm = !showForm}
			class="bg-pink-400 hover:bg-pink-500 text-white text-sm px-4 py-2 rounded-full transition-colors shadow-md"
		>
			{showForm ? 'Cancel' : '+ Write'}
		</button>
	</div>

	{#if showForm}
		<form onsubmit={handleSubmit} class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 space-y-3">
			<input
				type="text"
				bind:value={title}
				placeholder="Title..."
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
			/>
			<textarea
				bind:value={content}
				placeholder="Write your heart out..."
				rows="4"
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30 resize-none"
			></textarea>
			<div class="grid grid-cols-2 gap-3">
				<select
					bind:value={noteType}
					class="px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
				>
					{#each Object.entries(typeLabels) as [value, label]}
						<option {value}>{label}</option>
					{/each}
				</select>
				<select
					bind:value={visibleTo}
					class="px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
				>
					<option value="both">Both</option>
					<option value="ryan">Ryan only</option>
					<option value="dira">Dira only</option>
				</select>
			</div>
			<button
				type="submit"
				disabled={submitting || !title.trim() || !content.trim()}
				class="w-full bg-pink-400 hover:bg-pink-500 disabled:bg-pink-200 text-white py-2.5 rounded-xl transition-colors text-sm font-medium"
			>
				{submitting ? 'Sending...' : 'Send Note 💕'}
			</button>
		</form>
	{/if}

	{#if loading}
		<div class="text-center text-pink-300 py-8 animate-pulse">Loading notes...</div>
	{:else if notes.length === 0}
		<div class="bg-white rounded-2xl shadow-md shadow-pink-50 p-8 border border-pink-100 text-center">
			<div class="text-4xl mb-3">💌</div>
			<p class="text-pink-400">No secret notes yet!</p>
			<p class="text-pink-300 text-sm mt-1">Write your first love note</p>
		</div>
	{:else}
		<div class="space-y-3">
			{#each notes as note}
				<div class="bg-white rounded-2xl shadow-md shadow-pink-50 p-4 border border-pink-100">
					<div class="flex items-start justify-between mb-2">
						<h3 class="font-medium text-pink-600 text-sm">{note.title}</h3>
						{#if note.note_type}
							<span class="text-xs bg-pink-50 text-pink-400 px-2 py-0.5 rounded-full flex-shrink-0 ml-2">
								{typeLabels[note.note_type] || note.note_type}
							</span>
						{/if}
					</div>
					<p class="text-pink-500/80 text-sm whitespace-pre-wrap leading-relaxed">{note.content}</p>
					<div class="flex items-center gap-2 mt-3">
						{#if note.visible_to}
							<span class="text-xs text-pink-300">
								{note.visible_to === 'both' ? '👀 Both' : note.visible_to === 'ryan' ? '👦 Ryan only' : '👧 Dira only'}
							</span>
						{/if}
						<span class="text-xs text-pink-200 ml-auto">
							{new Date(note.created_at).toLocaleDateString()}
						</span>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>

<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';

	let items: any[] = $state([]);
	let loading = $state(true);
	let showForm = $state(false);
	let title = $state('');
	let caption = $state('');
	let category = $state('moment');
	let takenAt = $state('');
	let imageFile: File | null = $state(null);
	let submitting = $state(false);

	onMount(async () => {
		await loadGallery();
	});

	async function loadGallery() {
		loading = true;
		try {
			items = await api.getGallery();
		} catch {
			items = [];
		} finally {
			loading = false;
		}
	}

	function handleFileSelect(e: Event) {
		const input = e.target as HTMLInputElement;
		if (input.files && input.files[0]) {
			imageFile = input.files[0];
		}
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		if (!imageFile) return;
		submitting = true;
		try {
			const formData = new FormData();
			formData.append('image', imageFile);
			formData.append('title', title.trim() || imageFile.name);
			if (caption.trim()) formData.append('caption', caption.trim());
			if (category) formData.append('category', category);
			if (takenAt) formData.append('taken_at', takenAt);

			await api.uploadImage(formData);
			title = '';
			caption = '';
			category = 'moment';
			takenAt = '';
			imageFile = null;
			showForm = false;
			await loadGallery();
		} catch {
		} finally {
			submitting = false;
		}
	}

	async function deleteItem(id: number) {
		if (!confirm('Delete this photo?')) return;
		try {
			await api.deleteGalleryItem(id);
			await loadGallery();
		} catch {}
	}
</script>

<svelte:head>
	<title>Gallery - diraaax</title>
</svelte:head>

<div class="space-y-4">
	<div class="flex items-center justify-between">
		<h1 class="text-xl font-bold text-pink-500">📸 Our Gallery</h1>
		<button
			onclick={() => showForm = !showForm}
			class="bg-pink-400 hover:bg-pink-500 text-white text-sm px-4 py-2 rounded-full transition-colors shadow-md"
		>
			{showForm ? 'Cancel' : '+ Upload'}
		</button>
	</div>

	{#if showForm}
		<form onsubmit={handleSubmit} class="bg-white rounded-2xl shadow-md shadow-pink-50 p-5 border border-pink-100 space-y-3">
			<label class="block">
				<span class="text-sm text-pink-400 mb-1 block">Photo</span>
				<input
					type="file"
					accept="image/*"
					onchange={handleFileSelect}
					class="w-full text-sm text-pink-400 file:mr-3 file:py-2 file:px-4 file:rounded-xl file:border-0 file:text-sm file:bg-pink-50 file:text-pink-500 hover:file:bg-pink-100"
				/>
			</label>
			<input
				type="text"
				bind:value={title}
				placeholder="Title (optional)"
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
			/>
			<textarea
				bind:value={caption}
				placeholder="Caption (optional)"
				rows="2"
				class="w-full px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30 resize-none"
			></textarea>
			<div class="grid grid-cols-2 gap-3">
				<select
					bind:value={category}
					class="px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
				>
					<option value="moment">Moment</option>
					<option value="selfie">Selfie</option>
					<option value="date">Date</option>
					<option value="travel">Travel</option>
					<option value="food">Food</option>
					<option value="other">Other</option>
				</select>
				<input
					type="date"
					bind:value={takenAt}
					class="px-4 py-2.5 rounded-xl border border-pink-100 focus:border-pink-300 focus:outline-none text-sm bg-pink-50/30"
				/>
			</div>
			<button
				type="submit"
				disabled={submitting || !imageFile}
				class="w-full bg-pink-400 hover:bg-pink-500 disabled:bg-pink-200 text-white py-2.5 rounded-xl transition-colors text-sm font-medium"
			>
				{submitting ? 'Uploading...' : 'Upload Photo 📸'}
			</button>
		</form>
	{/if}

	{#if loading}
		<div class="text-center text-pink-300 py-8 animate-pulse">Loading gallery...</div>
	{:else if items.length === 0}
		<div class="bg-white rounded-2xl shadow-md shadow-pink-50 p-8 border border-pink-100 text-center">
			<div class="text-4xl mb-3">📷</div>
			<p class="text-pink-400">No photos yet!</p>
			<p class="text-pink-300 text-sm mt-1">Upload your first memory</p>
		</div>
	{:else}
		<div class="grid grid-cols-2 gap-3">
			{#each items as item}
				<div class="bg-white rounded-2xl shadow-md shadow-pink-50 border border-pink-100 overflow-hidden group">
					<div class="aspect-square overflow-hidden">
						<img
							src={api.imageUrl(item.file_name)}
							alt={item.title}
							class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
						/>
					</div>
					<div class="p-3">
						<h3 class="font-medium text-pink-600 text-xs truncate">{item.title}</h3>
						{#if item.caption}
							<p class="text-pink-300 text-xs mt-0.5 truncate">{item.caption}</p>
						{/if}
						<div class="flex items-center justify-between mt-2">
							{#if item.category}
								<span class="text-xs bg-pink-50 text-pink-400 px-2 py-0.5 rounded-full">{item.category}</span>
							{/if}
							<button
								onclick={() => deleteItem(item.id)}
								class="text-pink-200 hover:text-red-400 transition-colors text-xs"
							>
								✕
							</button>
						</div>
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>

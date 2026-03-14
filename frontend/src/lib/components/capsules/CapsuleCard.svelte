<script lang="ts">
	import type { LoveCapsule } from '$lib/types';
	import { formatDate, canOpenCapsule } from '$lib/utils';

	let {
		capsule,
		ondelete
	}: {
		capsule: LoveCapsule;
		ondelete?: (id: number) => void;
	} = $props();

		const isOpened = $derived(Boolean(capsule.is_opened));
		const canOpen = $derived(canOpenCapsule(capsule.open_date));
</script>

<a
	href="/capsules/{capsule.id}"
	class="love-card block p-5 md:p-6 relative overflow-hidden group {!isOpened && !canOpen ? 'opacity-75' : ''}"
>
	<!-- Decorative seal -->
	{#if !isOpened}
		<div class="absolute top-3 right-3">
			{#if canOpen}
				<span class="inline-flex items-center gap-1 text-xs font-semibold bg-pink-500 text-white px-3 py-1 rounded-full animate-pulse-soft">
					💌 Buka Sekarang!
				</span>
			{:else}
				<span class="inline-flex items-center gap-1 text-xs font-semibold bg-pink-100 text-pink-400 px-3 py-1 rounded-full">
					🔒 Terkunci
				</span>
			{/if}
		</div>
	{:else}
		<div class="absolute top-3 right-3">
			<span class="inline-flex items-center gap-1 text-xs font-semibold bg-green-100 text-green-600 px-3 py-1 rounded-full">
				✅ Terbuka
			</span>
		</div>
	{/if}

	<div class="flex items-start gap-4 pr-20 md:pr-24">
		<div class="text-3xl shrink-0">
			{#if isOpened}
				💝
			{:else if canOpen}
				💌
			{:else}
				🔐
			{/if}
		</div>

		<div class="min-w-0">
			<h3 class="font-bold text-rose-deep text-base mb-1 truncate">{capsule.title}</h3>
			<p class="text-rose-muted text-xs">
				{#if !isOpened}
					Dibuka pada: {formatDate(capsule.open_date, 'short')}
				{:else}
					Sudah dibuka 💖
				{/if}
			</p>
		</div>
	</div>

	{#if ondelete}
		<div class="mt-3 pt-3 border-t border-pink-100/80">
			<button
				type="button"
				onclick={(event) => {
					event.preventDefault();
					event.stopPropagation();
					ondelete?.(capsule.id);
				}}
				class="text-xs text-red-400 hover:text-red-500"
			>
				Hapus Capsule
			</button>
		</div>
	{/if}

	<!-- Bottom glow effect -->
	{#if canOpen && !isOpened}
		<div class="absolute bottom-0 left-0 right-0 h-1 bg-linear-to-r from-pink-400 via-pink-500 to-pink-400"></div>
	{/if}
</a>

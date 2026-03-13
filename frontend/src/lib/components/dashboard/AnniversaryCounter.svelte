<script lang="ts">
	import { getAnniversaryInfo } from '$lib/utils';

	// Anniversary date: 12 December 2025
	const ANNIVERSARY = '2025-12-12';
	const info = getAnniversaryInfo(ANNIVERSARY);
</script>

<div class="love-card-static p-5 md:p-6 text-center bg-gradient-to-br from-white to-pink-50">
	<div class="text-3xl mb-3 animate-pulse-soft">💍</div>
	<p class="text-rose-muted text-xs uppercase tracking-wider mb-2">Kita sudah bersama</p>

	<div class="flex items-center justify-center gap-3 mb-3">
		<div class="text-center">
			<div class="text-2xl md:text-3xl font-extrabold text-pink-500">{info.totalDays}</div>
			<div class="text-xs text-rose-muted">Hari</div>
		</div>
	</div>

	{#if info.years > 0 || info.months > 0}
		<p class="text-pink-400 text-sm font-medium">
			{#if info.years > 0}{info.years} tahun{/if}
			{#if info.months > 0} {info.months} bulan{/if}
			{#if info.days > 0} {info.days} hari{/if}
		</p>
	{/if}

	<p class="text-rose-muted text-xs mt-2">Sejak {info.formattedDate}</p>
	<p class="text-pink-300 text-xs mt-1">...dan selamanya 💕</p>
</div>

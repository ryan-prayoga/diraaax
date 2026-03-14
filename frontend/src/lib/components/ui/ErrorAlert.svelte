<script lang="ts">
	let {
		message = 'Terjadi kesalahan',
		class: className = ''
	}: {
		message?: string;
		class?: string;
	} = $props();
</script>

<div
	role="alert"
	class="rounded-2xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-500 {className}"
>
	{message}
</div>

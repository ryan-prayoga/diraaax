<script lang="ts">
	let {
		class: className = '',
		hover = true,
		children
	}: {
		class?: string;
		hover?: boolean;
		children: any;
	} = $props();
</script>

<div class="{hover ? 'love-card' : 'love-card-static'} p-5 md:p-6 {className}">
	{@render children()}
</div>

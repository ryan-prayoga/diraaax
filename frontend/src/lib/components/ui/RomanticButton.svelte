<script lang="ts">
	let {
		variant = 'primary',
		size = 'md',
		class: className = '',
		disabled = false,
		type = 'button',
		onclick,
		children,
		...restProps
	}: {
		variant?: 'primary' | 'secondary' | 'ghost';
		size?: 'sm' | 'md' | 'lg';
		class?: string;
		disabled?: boolean;
		type?: 'button' | 'submit' | 'reset';
		onclick?: (e: MouseEvent) => void;
		children: any;
		[key: string]: any;
	} = $props();

	const sizeClasses = {
		sm: 'text-sm px-4 py-2',
		md: 'px-6 py-3',
		lg: 'text-lg px-8 py-4'
	};

	const variantClasses = {
		primary: 'btn-primary',
		secondary: 'btn-secondary',
		ghost: 'bg-transparent text-pink-500 hover:bg-pink-50 rounded-xl transition-colors px-4 py-2 font-medium'
	};
</script>

<button
	{type}
	{disabled}
	{onclick}
	class="{variantClasses[variant]} {variant !== 'ghost' ? sizeClasses[size] : ''} {className}"
	{...restProps}
>
	{@render children()}
</button>

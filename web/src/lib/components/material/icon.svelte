<script lang="ts">
	import type { MaterialSymbolToken } from '$lib/types/material/IconEnum';

	export let token: MaterialSymbolToken;
	export let style: 'rounded' | 'outlined' | 'sharp' = 'rounded';
</script>

<span class="material-symbols-{style}">{token}</span>

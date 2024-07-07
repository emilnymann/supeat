<script lang="ts">
	import ContentContainer from '$lib/components/layout/content_container.svelte';
	import Icon from '$lib/components/material/icon.svelte';
	import { MaterialSymbolToken as Icons } from '$lib/types/material/IconEnum';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<h1>Frokost</h1>

<div class="flex flex-col gap-4">
	{#each data.days as day}
		<ContentContainer>
			<header>
				<h2>{day.title}</h2>
			</header>

			<article class="main-dish-container">
				<header>
					<Icon token={Icons.DinnerDining} />
					<h3>{day.main.title}</h3>
				</header>
				<p>{day.main.description}</p>
			</article>

			<div class="side-dishes-container">
				<article class="card">
					<header>
						<Icon token={Icons.BreakfastDining} />
						<h3>{day.bread.title}</h3>
					</header>
					<p>{day.bread.description}</p>
				</article>

				<article class="card">
					<header>
						<Icon token={Icons.EggAlt} />
						<h3>{day.toppings.title}</h3>
					</header>
					<p>{day.toppings.description}</p>
				</article>

				<article class="card">
					<header>
						<Icon token={Icons.Nutrition} />
						<h3>{day.salads.title}</h3>
					</header>
					<p>{day.salads.description}</p>
				</article>
			</div>
		</ContentContainer>
	{/each}
</div>

<style lang="postcss">
	h1 {
		@apply text-4xl;
		@apply p-4;
	}

	h2 {
		@apply text-2xl;
	}

	.card {
		@apply bg-surface-container-high;
		@apply p-4;
		@apply rounded-xl;
	}

	.main-dish-container header,
	.card header {
		@apply text-xl;
		display: flex;
		align-items: center;
		@apply gap-2;
		@apply mb-2;
	}

	.side-dishes-container {
		@apply flex;
		@apply flex-col;
		@apply md:flex-row;
		@apply gap-2;

		& .card {
			@apply flex-grow;
		}
	}
</style>

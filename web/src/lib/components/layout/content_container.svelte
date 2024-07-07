<section>
	<slot />
</section>

<style lang="postcss">
	section {
		@apply bg-surface;
		@apply rounded-3xl;
		@apply p-4;
		display: flex;
		flex-direction: column;
		@apply gap-4;
	}
</style>

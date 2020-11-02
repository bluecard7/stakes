<script>
	import ClockButton from './ClockButton.svelte';
	let user = 'Paul';
</script>

<main> 
	<h3>Hi {user}</h3>
	<ClockButton />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h3 {
		color: #ff3e00;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
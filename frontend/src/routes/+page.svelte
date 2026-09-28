<script lang="ts">
	let message = $state('');

	async function hello() {
		const response = await fetch('http://localhost:8080/api/hello');
		const data = await response.json();

		message = data.message;
	}
</script>

<button onclick={hello}>Hello</button>

<p>{message}</p>
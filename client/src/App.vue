<script setup lang="ts">
import { TonConnectUIProvider } from '@townsquarelabs/ui-vue';

import MainView from './views/MainView.vue';

const options = {
	manifestUrl: 'https://redacted.example.com/tonconnect-manifest.json',
};
</script>

<template>
	<TonConnectUIProvider :options="options">
		<Suspense>
			<MainView />
		</Suspense>
	</TonConnectUIProvider>
</template>

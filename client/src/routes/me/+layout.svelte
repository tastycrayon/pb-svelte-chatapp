<script lang="ts">
	import Icon from '$components/icon.svelte';
	import { PROFILE_EDIT_PATH, ROOM_PATH } from '$lib/constant';
	import { AppBar, AppShell } from '@skeletonlabs/skeleton';
</script>

<AppShell>
	<svelte:fragment slot="header">
		<AppBar>
			<svelte:fragment slot="lead">
				<a href={ROOM_PATH} class="btn-icon btn-icon-sm variant-soft">
					<Icon name="arrow-left" width="18px" height="18px" />
				</a>
			</svelte:fragment>
			<a href={ROOM_PATH}>Back to rooms</a>
			<svelte:fragment slot="trail">
				<a class="btn-icon btn-icon-sm variant-soft" href={PROFILE_EDIT_PATH} title="Edit Profile">
					<Icon name="edit" width="18px" height="18px" />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<!-- (sidebarLeft) -->
	<!-- (sidebarRight) -->
	<!-- (pageHeader) -->
	<!-- Router Slot -->
	<slot />
	<!-- ---- / ---- -->
	<!-- (pageFooter) -->
	<!-- (footer) -->
</AppShell>

<template>
	<view>
		<u-tabbar :value="tabeIndex" @change="name => tabeIndex = name" :fixed="true" :placeholder="true"
			:safeAreaInsetBottom="true">
			<u-tabbar-item text="首页" icon="home" @click="navigate"></u-tabbar-item>
			<u-tabbar-item text="创作者" icon="photo" @click="navigate"></u-tabbar-item>
			<u-tabbar-item text="我的" icon="account" @click="navigate"></u-tabbar-item>
		</u-tabbar>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				tabeIndex:2,
			};
		},
		methods:{
			navigate(){
				const urlArray=['/pages/index/index','/pages/author/author','/pages/me/me']
				uni.navigateTo({
					url: urlArray[tabeIndex],
				});
			}
		}
	}
</script>

<style lang="scss">

</style>

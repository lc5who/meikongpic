<template>
	<view class="author_avatar_box">
		<image :src="src" class="author_avatar_image"></image>
	</view>
</template>

<script>
	export default {
		name:"avatar",
		props:{
			src: {
				type: String,
				default: ''
			}
		},
		data() {
			return {
				
			};
		}
	}
</script>

<style lang="scss">
	.author_avatar_box {
		display: inline-block;
		width: 100upx;
		height: 100upx;
	}
	
	.author_avatar_image {
		width: 100upx;
		height: 100upx;
		border-radius: 50%;
		border: solid 5upx white;
	}
</style>
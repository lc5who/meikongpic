<template>
	<view class="container">
		<view class="back">
			<navigator hover-class="navigator-hover" open-type="navigateBack">
				<uni-icons type="back" size="30"></uni-icons>
			</navigator>
		</view>
		<view class="content">
			<template v-if="downRecord.length!=0">
				<navigator :url="'/pages/index/detail/detail?id='+item.ID" v-for="(item,index) in downRecord" :key="index">
					<view class="content_image">
						<image :src="item.url" mode=""></image>
					</view>
				</navigator>
			</template>
		</view>
	</view>
</template>

<script setup>
	
	import {onMounted,ref} from 'vue'
	let downRecord = ref([])
	onMounted(()=>{
		let dataStr = localStorage.getItem('download')
		let data = JSON.parse(dataStr)
		downRecord.value = data
		console.log(downRecord.value);
	})
</script>

<style lang="scss">
	.container {
		width: 100vw;
		height: 100vh;
		background-color: black;
		overflow-y: auto;
		overflow-x: hidden;
	}
	.content {
		padding: 20upx;
		display: grid;
		grid-template-columns: repeat(3, minmax(0, 1fr));
		gap: 20upx;
	}
	.content_image {
		height: 400upx;
	}
	
	.content_image image {
		height: 100%;
		width: 100%;
		border: 1px solid #fff;
		border-radius: 10px;
	
	}
	.back{
	
		// position: fixed;
		width: 60upx;
		height: 60upx;
		margin-top: 20upx;
		margin-left: 20upx;
		// top: 40upx;
		// left: 30upx;
		z-index: 11;
		border: 1px solid white;
		border-radius: 15px;
		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>

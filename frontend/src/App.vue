<template>
	<el-config-provider :locale="zhCn">
		<router-view></router-view>
	</el-config-provider>
</template>

<script>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
export default {
	data() {
		return {
			zhCn: zhCn
		}
	},
}
</script>

<style lang="less">
body {
	margin: 0;
	padding: 0;
}
</style>

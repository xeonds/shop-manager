<template>
  <div class="customer-view">
    <el-row>
      <el-col :span="24">
        <h1>Dashboard</h1>
        <p>Welcome. Get prepared for today's work?</p>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24">
        <el-card>
          <template #header>
            <el-text>Analysis</el-text>
          </template>
          <el-row>
            <el-col :span="6">
              <el-statistic title="All customers" :value="customerList.length" />
            </el-col>
            <el-col :span="6">
              <el-statistic title="Ratio of men to women" :value="men2Women"></el-statistic>
            </el-col>
            <el-col :span="6">
              <el-statistic title="Premier Customers" :value="premierCustomers" />
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import customerAPI from '../../api/customer'

export default {
  data() {
    return {
      customer: {
        name: '',
        email: '',
        phone: ''
      },
      customerList: [
        { id: 1, name: 'John Doe', email: 'johndoe@example.com', phone: '1234567890', premier: true, sex: "man" },
        { id: 2, name: 'Jane Doe', email: 'janedoe@example.com', phone: '0987654321', premier: true, sex: "woman" }
      ],
    }
  },
  computed: {
    filteredCustomerList() {
      return this.customerList.filter(customer => {
        return customer.name.toLowerCase().includes(this.search.toLowerCase())
          || customer.email.toLowerCase().includes(this.search.toLowerCase())
          || customer.phone.toLowerCase().includes(this.search.toLowerCase())
          || customer.id.toString().includes(this.search)
      })
    },
    men2Women() {
      return this.customerList.filter(customer => { customer.sex == 'man' }).length
        + '/' + this.customerList.filter(customer => { customer.sex == 'man' }).length
    },
    premierCustomers() {
      return this.customerList.filter(customer => { customer.premier == true }).length
    }
  },
  methods: {
    async showCustomerList() {
      try {
        const response = await customerAPI.getCustomers()
        this.customerList = response.data
      } catch (error) {
        this.$message.error('Get customer list failed')
      }
    },
  },
  mounted() {
    this.showCustomerList()
  }
}
</script>

<style lang="less" scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.customer-view {
  max-width: 1000px;
  margin: 0 auto;

  .el-row {
    margin-bottom: 2rem;
  }
}
</style>


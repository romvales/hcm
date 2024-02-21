import { test, afterEach, expect } from 'bun:test'
import { Supabase_HCMWorkerService } from './HCMWorkerService'
import { Supabase_HCMDeductionService } from './HCMDeductionService'
import { createClient } from '@supabase/supabase-js'
import { Database, DeductionScope, DeductionStatus, DeductionType } from '../../../src'
import { createRandomWorker, setupFakeWorkerFields } from './HCMWorkerService.spec'

const localClient = createClient<Database>(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_SERVICE_ROLE_KEY as string,
)

const LOCALCLIENT_TEST_EMAIL = 'test@romvales.com'
const LOCALCLIENT_TEST_PASSWORD = 'testpassword'

let userId: string

afterEach(async () => {
  await cleanHCMDeductionServiceTest()
})

const setupHCMDeductionServiceTest = async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const deductionService = new Supabase_HCMDeductionService(localClient, workerService)

  await localClient.auth.signUp({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })
  

  const { data, error } = await localClient.auth.signInWithPassword({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { user } = data
  
  if (error) throw error
  if (user) {
    userId = user.id
  }

  const _worker = createRandomWorker()
  const worker = await setupFakeWorkerFields(
    workerService,
    workerService.createWorker(_worker.email, _worker.username),
    user,
    _worker
  ).saveWorker()


  return {
    deductionService,
    workerService,
    
    user,
    _worker,
    worker,
  }
}

const cleanHCMDeductionServiceTest = async () => {
  await localClient.auth.signOut()
  if (userId) await localClient.auth.admin.deleteUser(userId)
}


test('> HCMDeductionService.createDeduction()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  expect(deduct).toBeDefined()
  expect(deduct.name).toBe(DEDUCTION_NAME)
  expect(deduct.value).toBe(DEDUCTION_VALUE)
})

test('> HCMDeductionService.getDeductionById()', async () => {
  const {} = await setupHCMDeductionServiceTest()

  throw Error()
})

test('> HCMDeductionService.deleteDeductionById()', async () => {
  const {} = await setupHCMDeductionServiceTest()

  throw Error()
})

test('> HCMDeductionService.saveDeduction()', async () => {
  const {} = await setupHCMDeductionServiceTest()

  throw Error()
})

test('> HCMDeductionService.changeType()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeType(DeductionType.BENEFIT)

  expect(deduct).toBeDefined()
  expect(deduct.type).toBe(DeductionType.BENEFIT)
})

test('> HCMDeductionService.changeScope()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeScope(DeductionScope.GLOBAL)

  expect(deduct.scope).toBe(DeductionScope.GLOBAL)
})

test('> HCMDeductionService.changeStatus()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeStatus(DeductionStatus.PENDING)

  expect(deduct.status).toBe(DeductionStatus.PENDING)
})

test('> HCMDeductionService.changeVoluntary()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeVoluntary(!deduct.isVoluntary)

  expect(deduct.isVoluntary).toBe(true)
})

test('> HCMDeductionService.changeName()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (3%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const NEW_NAME = 'SSS Contribution (3%)'

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeName(NEW_NAME)

  expect(deduct.name).toBe(NEW_NAME)
})

test('> HCMDeductionService.changeValue()', async () => {
  const { deductionService } = await setupHCMDeductionServiceTest()

  const DEDUCTION_NAME = 'SSS Monthly Contribution (5%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const NEW_VALUE = 15_000*0.05

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.changeValue(NEW_VALUE)

  expect(deduct.value).toBe(NEW_VALUE)
})

test('> HCMDeductionService.assignDeductionToWorker()', async () => {
  const { deductionService, worker } = await setupHCMDeductionServiceTest()


  const DEDUCTION_NAME = 'SSS Monthly Contribution (5%)'
  const DEDUCTION_VALUE = 15_000*0.03

  const deduct = deductionService.createDeduction(DEDUCTION_NAME, DEDUCTION_VALUE)

  deductionService.assignDeductionToWorker(worker)

  expect<number | undefined>(deduct.workerId).toBe(worker.id)
})
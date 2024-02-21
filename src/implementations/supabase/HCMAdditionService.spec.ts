import { createClient } from '@supabase/supabase-js'
import { expect, test, afterEach } from 'bun:test'
import { Database, AdditionScope, AdditionStatus, AdditionType } from '../../../src'

import { Supabase_HCMAdditionService } from './HCMAdditionService'
import { Supabase_HCMWorkerService } from './HCMWorkerService'
import { createRandomWorker, setupFakeWorkerFields } from './HCMWorkerService.spec'

const localClient = createClient<Database>(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_SERVICE_ROLE_KEY as string,
)

const LOCALCLIENT_TEST_EMAIL = 'test@romvales.com'
const LOCALCLIENT_TEST_PASSWORD = 'testpassword'

let userId: string

afterEach(async () => {
  await cleanHCMAdditionServiceTest(userId)
})

const setupHCMAdditionServiceTest = async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const additionService = new Supabase_HCMAdditionService(localClient, workerService)

  await localClient.auth.signUp({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { data, error } = await localClient.auth.signInWithPassword({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { user } = data

  if (user?.id) userId = user.id
  if (error) throw error

  const _worker = createRandomWorker()
  const _createdWorker = workerService.createWorker(_worker.email, _worker.username)

  const worker = await setupFakeWorkerFields(workerService, _createdWorker, user, _worker).saveWorker()

  return {
    additionService,
    workerService,

    user,

    _createdWorker,
    worker,
  }
}

const cleanHCMAdditionServiceTest = async (userId?: string) => {
  await localClient.auth.signOut()
  if (userId) await localClient.auth.admin.deleteUser(userId)

  userId = undefined
}

test('> HCMAdditionService.createAddition()', async () => {
  const ADDED_NAME = 'Mandatory Contribution'
  const ADDED_VALUE = 100.00

  const { additionService } = await setupHCMAdditionServiceTest()
  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  expect(added).toBeDefined()
  expect(added.value).toBe(ADDED_VALUE)
  expect(added.name).toBe(ADDED_NAME)

})

test('> HCMAdditionService.getAdditionById()', async () => {
  await setupHCMAdditionServiceTest()

  throw Error()
})

test('> HCMAdditionService.deleteAdditionById()', async () => {
  await setupHCMAdditionServiceTest()

  throw Error()
})

test('> HCMAdditionService.saveAddition()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'Employee of the Month Bonus'
  const ADDED_VALUE = 100.00

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)
  const savedAdded = await additionService.saveAddition()

  expect(savedAdded).toBeDefined()
  expect(savedAdded.id).toBeDefined()
  expect<string | undefined>(savedAdded.name).toBe(added.name)

  await localClient.from('additions').delete().match({ id: savedAdded.id }).throwOnError()
})

test('> HCMAdditionService.changeValue()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'Pag-ibig Monthly Contribution (Employee Share) Bonus Pkg.'
  const ADDED_VALUE = 100.00
  const NEW_VALUE = 15_000*0.03

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)
  
  additionService.changeValue(NEW_VALUE)

  expect(added.value).toBe(NEW_VALUE)
})

test('> HCMAdditionService.changeName()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'SSS Monthly Contribution (EE) Bonus Pkg.'
  const ADDED_VALUE = 15_000 * 0.07
  const NEW_NAME = 'SSS Monthly Contribution (ER) Bonus Pkg.'

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService.changeName(NEW_NAME)

  expect(added.name).toBe(NEW_NAME)
})

test('> HCMAdditionService.assignAdditionToWorker()', async () => {
  const { additionService, worker } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = '0.25% Added High Performance Bonus+'
  const ADDED_VALUE = 15_000*0.0025

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService.assignAdditionToWorker(worker)

  expect(added.workerId).toBeDefined()
  expect<number | undefined>(added.workerId).toBe(worker.id)
})

test('> HCMAdditionService.changeType()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'SSS Monthly Contribution (EE) Bonus Pkg.'
  const ADDED_VALUE = 15_000 * 0.07

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService.changeType(AdditionType.BONUS)

  expect(added.type).toBe(AdditionType.BONUS)
})

test('> HCMAdditionService.changeScope()', async () => {
  const { additionService, worker } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'SSS Monthly Contribution (EE) Bonus Pkg.'
  const ADDED_VALUE = 15_000 * 0.07

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService
    .changeScope(AdditionScope.WORKER)
    .assignAdditionToWorker(worker)

  expect(added.scope).toBe(AdditionScope.WORKER)
  expect<number | undefined>(added.workerId).toBe(worker.id)
})

test('> HCMAdditionService.changeStatus()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'SSS Monthly Contribution (EE) Bonus Pkg.'
  const ADDED_VALUE = 15_000 * 0.07

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService
    .changeStatus(AdditionStatus.PENDING)

  expect(added.status).toBe(AdditionStatus.PENDING)

})

test('> HCMAdditionService.setEphemeral()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'SSS Monthly Contribution (EE) Bonus Pkg.'
  const ADDED_VALUE = 15_000 * 0.07

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  additionService.setEphemeral(true)

  expect(added.isEphemeral).toBeTrue()
})

test('> HCMAdditionService.changeEffectiveDate()', async () => {
  const { additionService } = await setupHCMAdditionServiceTest()

  const ADDED_NAME = 'PhilHealth Monthly Contribution (3%)'
  const ADDED_VALUE = 15_000 * 0.03
  const EFFECTIVE_DATE = new Date().toJSON()

  const added = additionService.createAddition(ADDED_NAME, ADDED_VALUE)

  
  additionService.changeEffectiveDate(EFFECTIVE_DATE)

  expect(added.effectiveAt).toBe(EFFECTIVE_DATE)
})
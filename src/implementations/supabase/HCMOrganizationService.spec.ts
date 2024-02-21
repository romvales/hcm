import { randCompanyName } from '@ngneat/falso'
import { afterEach, expect, test } from 'bun:test'
import { Supabase_HCMOrganizationService } from './HCMOrganizationService'
import { User, createClient } from '@supabase/supabase-js'
import { Database, OrganizationIndustry, OrganizationStatus } from '../../../src'
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
  await cleanHCMOrganizationServiceTest(userId)
})

const setupHCMOrganizationServiceTest = async () => {
  await localClient.auth.signUp({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const { data } = await localClient.auth.signInWithPassword({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const user = data.user as User
  const oldOrgName = randCompanyName()
  const newOrgName = randCompanyName()

  if (user) userId = user.id

  const workerService = new Supabase_HCMWorkerService(localClient)
  const orgService = new Supabase_HCMOrganizationService(localClient, workerService)

  const _orgCreator = createRandomWorker()
  let orgCreator = workerService.createWorker(_orgCreator.email, _orgCreator.username)

  orgCreator = await setupFakeWorkerFields(workerService, orgCreator, user, _orgCreator).saveWorker()

  return {
    oldOrgName,
    newOrgName,
    orgService,
    _orgCreator,
    orgCreator,
    workerService,
    user,
  }
}

const cleanHCMOrganizationServiceTest = async (id: string) => {
  await localClient.auth.signOut()
  if (id) await localClient.auth.admin.deleteUser(id)
}

test('> HCMOrganizationService.getOrgCreator()', async () => {
  const { 
    orgCreator,
    orgService,
    oldOrgName } = await setupHCMOrganizationServiceTest()

  orgService.createOrg(oldOrgName)

  const org = await orgService.saveOrg()
  const savedOrgCreator = await orgService.getOrgCreator()
    
  expect(savedOrgCreator).toBeDefined()
  expect(savedOrgCreator?.id).toBe(orgCreator.id as number)
  expect(savedOrgCreator?.email).toBe(orgCreator.email)

  await localClient.from('organizations').delete().match({ id: org.id })
})

test('> HCMOrganizationService.getOrgById()', async () => {
  const { 
    orgService, 
    oldOrgName } = await setupHCMOrganizationServiceTest()

  const unsavedOrg = orgService.createOrg(oldOrgName)
  const savedOrg = await orgService.setTarget(unsavedOrg)
    .changeOrgIndustry(OrganizationIndustry.EQUIPMENT)
    .changeOrgStatus(OrganizationStatus.INACTIVE)
    .saveOrg()

  const getSavedOrg = await orgService.getOrgById(savedOrg.id as number)

  expect(getSavedOrg).toBeDefined()

  await orgService.deleteOrgById(savedOrg.id)
})

test('> HCMOrganizationService.deleteOrgById()', async () => {
  const { 
    orgService,
    oldOrgName } = await setupHCMOrganizationServiceTest()

  const unsavedOrg = orgService.createOrg(oldOrgName)
  const org = await orgService.setTarget(unsavedOrg)
    .changeOrgIndustry(OrganizationIndustry.HOSPITALITY)
    .changeOrgStatus(OrganizationStatus.INACTIVE)
    .saveOrg()

  expect(org).toBeDefined()
  expect(org.id).toBeDefined()

  const delRes = await orgService.deleteOrgById(org.id)

  expect(delRes.error).toBeNull()
})

test('> HCMOrganizationService.changeOrgName()', async () => {
  const { orgService, oldOrgName, newOrgName } = await setupHCMOrganizationServiceTest()
  const org = orgService.createOrg(oldOrgName)

  expect(org).toBeDefined()
  expect(org.name).toBe(oldOrgName)

  orgService.changeOrgName(newOrgName) // Change the name of the organization

  expect(org.name).toBe(newOrgName)
})

test('> HCMOrganizationService.changeOrgIndustry()', async () => {
  const { orgService, oldOrgName } = await setupHCMOrganizationServiceTest()
  const org = orgService.createOrg(oldOrgName)

  expect(org).toBeDefined()
  expect(org.name).toBe(oldOrgName)

  orgService
    .changeOrgIndustry(OrganizationIndustry.HOSPITALITY)

  expect(org.industry).toBe(OrganizationIndustry.HOSPITALITY)
})

test('> HCMOrganizationService.changeOrgStatus()', async () => {
  const { orgService, oldOrgName } = await setupHCMOrganizationServiceTest()
  const org = orgService.createOrg(oldOrgName)

  expect(org).toBeDefined()
  expect(org.name).toBe(oldOrgName)
})

test('> HCMOrganizationService.saveOrg()', async () => {
  const { 
    orgService, 
    oldOrgName,
    orgCreator } = await setupHCMOrganizationServiceTest()

  const org = orgService.createOrg(oldOrgName)

  const savedOrg = await orgService
    .setTarget(org)
    .changeOrgIndustry(OrganizationIndustry.RETAIL)
    .changeOrgStatus(OrganizationStatus.INACTIVE)
    .saveOrg()

  expect(savedOrg).toBeDefined()
  expect(savedOrg.id).toBeDefined()
  expect(savedOrg.createdById).toBe(orgCreator.id as number)

  const savedOrg_orgCreator = await orgService.getOrgCreator()

  expect(savedOrg_orgCreator).toBeDefined()
  expect(savedOrg_orgCreator?.email).toBe(orgCreator.email)

  await orgService.deleteOrgById(savedOrg.id)
})

test('> HCMOrganizationService.removeWorkerFromOrgById()', async () => {
  await setupHCMOrganizationServiceTest()
  
})

test('> HCMOrganizationService.addWorkerToOrgById()', async () => {
  await setupHCMOrganizationServiceTest()

})

/* eslint-disable @typescript-eslint/no-var-requires */
/* eslint-disable no-undef */
const path = require('path')
const { Project } = require('ts-morph')

const GENERATED = 'src/lib/pb'

const fixImport = async dir => {
  const project = new Project()
  project.addSourceFilesAtPaths(`${dir}/**/*.js`)

  const sourceFile = project.getSourceFileOrThrow('message.js')
  const imports = sourceFile.getImportDeclarations()

  const im = imports.find(
    im => im.getModuleSpecifierValue() === 'protobufjs/minimal'
  )
  const importName = im.getNamespaceImport().getText()
  im.removeNamespaceImport()
  im.setDefaultImport(importName)

  await sourceFile.save()
}

;(async () => {
  await fixImport(path.resolve(__dirname, '../', GENERATED))
})()

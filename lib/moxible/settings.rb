module Moxible
  class Settings

    def self.createQuest(parsedOptions)
      quest = Moxible::Defaults::createDefaultQuest
      quest.merge! parsedOptions
      return quest
    end
  end
end

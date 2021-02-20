from googletrans import Translator
translator = Translator(service_urls=[
      'translate.google.com',
      'translate.google.co.kr',
    ])
translator.translate('안녕하세요.')
